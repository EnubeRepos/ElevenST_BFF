package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/EnubeRepos/ElevenST_BFF/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Log           *logrus.Logger
	Authorization string
	APIAuth       string
	Clinics       model.ClinicReferenceList
}

func New() (*Config, error) {
	var cfg Config

	// header := map[string]string{
	// 	"Method": "GET",
	// }
	var ClinicReferenceList = GetConfigurationFile()
	var AuthorizationAPI = os.Getenv("API_AUTH")
	// _, err := web.MakeRequest(ClinicReferenceList, header, nil, &clinicRef)
	// if err != nil {
	// 	return &cfg, err
	// }

	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   time.RFC3339Nano,
		DisableHTMLEscape: true,
		PrettyPrint:       true,
	})

	cfg.Authorization = os.Getenv("AUTH")
	cfg.Clinics = ClinicReferenceList
	cfg.APIAuth = AuthorizationAPI
	cfg.Log = l

	return &cfg, nil
}

func GetConfigurationFile() model.ClinicReferenceList {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("https://nyc3.digitaloceanspaces.com"),
		Region:   aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			"DO00VMQDDL8RNCRZMVCE",
			"X9nptrOAm/MZcjsGn0mmb4aTNb0jOhfVNZBhuJ+C9f4",
			""),
	})
	var clinics model.ClinicReferenceList
	var tempClinics model.ClinicReferenceList

	if err != nil {
		fmt.Println("Erro ao criar sessão AWS:", err)
		return clinics
	}

	svc := s3.New(sess)
	bucketName := "enube"
	objectKey := "ElevenST_BFF/clinicReference.json"
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	resp, err := svc.GetObject(params)

	if err != nil {
		fmt.Println("err:", err)
		return clinics
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err:", err)
		return clinics
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
	err = json.Unmarshal(body, &tempClinics)
	if err != nil {
		fmt.Println("err:", err)
		return clinics
	}
	for _, clinic := range tempClinics.List {
		if clinic.Active {
			clinics.List = append(clinics.List, clinic)
		}
	}
	return clinics
}
