package example

import (
	"encoding/json"
	"github.com/golfz/gorealtime"
	"github.com/stretchr/testify/assert"

	"math/rand"
	"testing"
	"time"
)

const (
	EddardStark = iota
	JonSnow
	DaenerysTargaryen
	JaimeLannister
	PetyrBaelish
)

func TestAMQPPublisher_Publish_Random_Device1(t *testing.T) {
	emp := getRandomEmployeeForTest()
	txn := getTransactionForTest("111111-1111", "1", "เครื่องที่ 1", emp)

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

func TestAMQPPublisher_Publish_Random_Device2(t *testing.T) {
	emp := getRandomEmployeeForTest()
	txn := getTransactionForTest("222222-2222", "2", "เครื่องที่ 2", emp)

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

// 11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111

func TestAMQPPublisher_Publish_Device1_EddardStark(t *testing.T) {
	txn := getTransactionForTest("111111-1111", "1", "เครื่องที่ 1", employees[EddardStark])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

func TestAMQPPublisher_Publish_Device1_JonSnow(t *testing.T) {
	txn := getTransactionForTest("111111-1111", "1", "เครื่องที่ 1", employees[JonSnow])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

func TestAMQPPublisher_Publish_Device1_DaenerysTargaryen(t *testing.T) {
	txn := getTransactionForTest("111111-1111", "1", "เครื่องที่ 1", employees[DaenerysTargaryen])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

func TestAMQPPublisher_Publish_Device1_JaimeLannister(t *testing.T) {
	txn := getTransactionForTest("111111-1111", "1", "เครื่องที่ 1", employees[JaimeLannister])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

func TestAMQPPublisher_Publish_Device1_PetyrBaelish(t *testing.T) {
	txn := getTransactionForTest("111111-1111", "1", "เครื่องที่ 1", employees[PetyrBaelish])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

// 22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222

func TestAMQPPublisher_Publish_Device2_EddardStark(t *testing.T) {
	txn := getTransactionForTest("222222-2222", "2", "เครื่องที่ 2", employees[EddardStark])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

func TestAMQPPublisher_Publish_Device2_JonSnow(t *testing.T) {
	txn := getTransactionForTest("222222-2222", "2", "เครื่องที่ 2", employees[JonSnow])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

func TestAMQPPublisher_Publish_Device2_DaenerysTargaryen(t *testing.T) {
	txn := getTransactionForTest("222222-2222", "2", "เครื่องที่ 2", employees[DaenerysTargaryen])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

func TestAMQPPublisher_Publish_Device2_JaimeLannister(t *testing.T) {
	txn := getTransactionForTest("222222-2222", "2", "เครื่องที่ 2", employees[JaimeLannister])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

func TestAMQPPublisher_Publish_Device2_PetyrBaelish(t *testing.T) {
	txn := getTransactionForTest("222222-2222", "2", "เครื่องที่ 2", employees[PetyrBaelish])

	bPayload, _ := json.Marshal(txn)

	err := gorealtime.NewAMQPPublisher("amqp://guest:guest@localhost:5672/realtime", "test").
		Publish(
			"transaction",
			string(bPayload),
			map[string]string{
				"employee_uuid": txn.Employee.EmployeeUUID,
				"hardware_uuid": txn.Hardware.HardwareUUID,
			},
		)
	assert.Nil(t, err)
}

// =================================================================================================

var employees = []Employee{
	{
		UUID:    "AAAAAA-111111",
		Code:    "111111",
		NameTH:  "เอ็ดดาร์ด สตาร์ค",
		NameEN:  "Eddard Stark",
		Picture: "https://thronesapi.com/assets/images/ned-stark.jpg",
	},
	{
		UUID:    "BBBBBB-222222",
		Code:    "222222",
		NameTH:  "จอน สโนว",
		NameEN:  "Jon Snow",
		Picture: "https://thronesapi.com/assets/images/jon-snow.jpg",
	},
	{
		UUID:    "CCCCCC-333333",
		Code:    "333333",
		NameTH:  "แดเนริส ทาแกเรียน",
		NameEN:  "Daenerys Targaryen",
		Picture: "https://thronesapi.com/assets/images/daenerys.jpg",
	},
	{
		UUID:    "DDDDDD-444444",
		Code:    "444444",
		NameTH:  "เจมี่ แลนนิสเตอร์",
		NameEN:  "Jaime Lannister",
		Picture: "https://thronesapi.com/assets/images/jaime-lannister.jpg",
	},
	{
		UUID:    "EEEEEE-555555",
		Code:    "555555",
		NameTH:  "ปีเตอร์ เบลิช",
		NameEN:  "Petyr Baelish",
		Picture: "https://thronesapi.com/assets/images/littlefinger.jpg",
	},
}

func getRandomEmployeeForTest() Employee {
	rand.Seed(time.Now().UnixNano())
	return employees[rand.Intn(len(employees))]
}

func getTransactionForTest(deviceUUID string, deviceCode string, titleTH string, emp Employee) TransactionMessage {
	tNow := time.Now()

	tStr := tNow.Format(time.RFC3339)

	return TransactionMessage{
		Company: CompanyObject{
			CompanyUUID: "AAAAAA-BBBB",
		},
		Employee: EmployeeObject{
			EmployeeUUID: emp.UUID,
			EmployeeCode: emp.Code,
			NameTH:       emp.NameTH,
			NameEN:       emp.NameEN,
		},
		Hardware: HardwareObject{
			HardwareUUID: deviceUUID,
			HardwareCode: deviceCode,
			TitleTH:      titleTH,
		},
		PictureURL: emp.Picture,
		TimeStamp:  tStr,
	}
}

// =================================================================================================

type Employee struct {
	UUID    string
	Code    string
	NameTH  string
	NameEN  string
	Picture string
}

type TransactionMessage struct {
	Company      CompanyObject      `json:"company"`
	Organization OrganizationObject `json:"organization"`
	Employee     EmployeeObject     `json:"employee"`
	Hardware     HardwareObject     `json:"hardware"`
	Location     LocationObject     `json:"location"`
	PictureURL   string             `json:"picture_url"`
	TimeStamp    string             `json:"time_stamp"`
}

type CompanyObject struct {
	CompanyUUID string `json:"company_uuid"`
	TitleTH     string `json:"title_th"`
	TitleEN     string `json:"title_en"`
}

type OrganizationObject struct {
	OrganizeUUID string `json:"organize_uuid"`
	TitleTH      string `json:"title_th"`
	TitleEn      string `json:"title_en"`
}

type EmployeeObject struct {
	EmployeeUUID string `json:"employee_uuid"`
	EmployeeCode string `json:"employee_code"`
	NameTH       string `json:"name_th"`
	NameEN       string `json:"name_en"`
}

type HardwareObject struct {
	HardwareUUID      string `json:"hardware_uuid"`
	HardwareCode      string `json:"hardware_code"`
	TitleTH           string `json:"title_th"`
	TitleEn           string `json:"title_en"`
	SerialNumber      string `json:"serial_number"`
	ModelUUID         string `json:"model_uuid"`
	ModelCode         string `json:"model_code"`
	ModelTitle        string `json:"model_title"`
	ManufacturerUUID  string `json:"manufacturer_uuid"`
	ManufacturerCode  string `json:"manufacturer_code"`
	ManufacturerTitle string `json:"manufacturer_title"`
}

type LocationObject struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	TitleTH   string `json:"title_th"`
	TitleEN   string `json:"title_en"`
}
