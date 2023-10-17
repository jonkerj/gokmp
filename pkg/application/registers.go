package application

import "fmt"

type (
	RegisterID uint16

	registerInfo struct {
		Id          RegisterID
		Name        string
		Description string
	}
)

const (
	RegisterHeatEnergy            = RegisterID(0x003c)
	RegisterFlowEnergy            = RegisterID(0x003d)
	RegisterReturnFlowEnergy      = RegisterID(0x003e)
	RegisterCoolingEnergy         = RegisterID(0x003f)
	RegisterTariffRegister2       = RegisterID(0x0040)
	RegisterTariffRegister3       = RegisterID(0x0041)
	RegisterTariffLimit2          = RegisterID(0x0042)
	RegisterTariffLimit3          = RegisterID(0x0043)
	RegisterVolumeRegister1       = RegisterID(0x0044)
	RegisterVolumeRegister2       = RegisterID(0x0045)
	RegisterMassRegister1         = RegisterID(0x0048)
	RegisterMassRegister2         = RegisterID(0x0049)
	RegisterCurrentInFlow         = RegisterID(0x004a)
	RegisterCurrentReturnFlow     = RegisterID(0x004b)
	RegisterCurrentPower          = RegisterID(0x0050)
	RegisterInputVA               = RegisterID(0x0054)
	RegisterInputVB               = RegisterID(0x0055)
	RegisterCurrentInTemp         = RegisterID(0x0056)
	RegisterCurrentReturnTemp     = RegisterID(0x0057)
	RegisterCurrentT3Temp         = RegisterID(0x0058)
	RegisterCurrentTempDifference = RegisterID(0x0059)
	RegisterCurrentInPressure     = RegisterID(0x005b)
	RegisterCurrentReturnPressure = RegisterID(0x005c)
	RegisterControlEnergy         = RegisterID(0x005e)
	RegisterTapEnergy             = RegisterID(0x005f)
	RegisterYEnergy               = RegisterID(0x0060)
	RegisterEnergy8               = RegisterID(0x0061)
	RegisterInfocode              = RegisterID(0x0063)
	RegisterMeterNoVB             = RegisterID(0x0068)
	RegisterEnergy9               = RegisterID(0x006e)
	RegisterMeterNo2              = RegisterID(0x0070)
	RegisterInfoEvent             = RegisterID(0x0071)
	RegisterMeterNoVA             = RegisterID(0x0072)
	RegisterCurrentT4Temp         = RegisterID(0x007a)
	RegisterMaxFlowYTDDate        = RegisterID(0x007b)
	RegisterMaxFlowYTDValue       = RegisterID(0x007c)
	RegisterMinFlowYTDDate        = RegisterID(0x007d)
	RegisterMinFlowYTDValue       = RegisterID(0x007e)
	RegisterMaxPowerYTDDate       = RegisterID(0x007f)
	RegisterMaxPowerYTDValue      = RegisterID(0x0080)
	RegisterMinPowerYTDDate       = RegisterID(0x0081)
	RegisterMinPowerYTDValue      = RegisterID(0x0082)
	RegisterMaxFlowMTDDate        = RegisterID(0x008a)
	RegisterMaxFlowMTDValue       = RegisterID(0x008b)
	RegisterMinFlowMTDDate        = RegisterID(0x008c)
	RegisterMinFlowMTDValue       = RegisterID(0x008d)
	RegisterMaxPowerMTDDate       = RegisterID(0x008e)
	RegisterMaxPowerMTDValue      = RegisterID(0x008f)
	RegisterMinPowerMTDDate       = RegisterID(0x0090)
	RegisterMinPowerMTDValue      = RegisterID(0x0091)
	RegisterAvgT1YTD              = RegisterID(0x0092)
	RegisterAvgT2YTD              = RegisterID(0x0093)
	RegisterAvgT1MTD              = RegisterID(0x0095)
	RegisterAvgT2MTD              = RegisterID(0x0096)
	RegisterConfigNo1             = RegisterID(0x0099)
	RegisterChecksum1             = RegisterID(0x009a)
	RegisterConfigNo2             = RegisterID(0x00a8)
	RegisterErrorHourCounter      = RegisterID(0x00af)
	RegisterInALiterImp           = RegisterID(0x00ea)
	RegisterInBLiterImp           = RegisterID(0x00eb)
	RegisterSerialNo              = RegisterID(0x03e9)
	RegisterClock                 = RegisterID(0x03ea)
	RegisterDate                  = RegisterID(0x03eb)
	RegisterOperationalHours      = RegisterID(0x03ec)
	RegisterSoftwareEdition       = RegisterID(0x03ed)
	RegisterMeterNo1              = RegisterID(0x03f2)
)

var regInfo = map[RegisterID]registerInfo{
	RegisterHeatEnergy: {
		Name:        "E1",
		Description: "Energy register 1: Heat energy",
	},
	RegisterFlowEnergy: {
		Name:        "E4",
		Description: "Energy register 4: Flow energy",
	},
	RegisterReturnFlowEnergy: {
		Name:        "E5",
		Description: "Energy register 5: Return flow energy",
	},
	RegisterCoolingEnergy: {
		Name:        "E3",
		Description: "Energy register 3: Cooling energy",
	},
	RegisterTariffRegister2: {
		Name:        "TA2",
		Description: "Tariff register 2",
	},
	RegisterTariffRegister3: {
		Name:        "TA3",
		Description: "Tariff register 3",
	},
	RegisterTariffLimit2: {
		Name:        "TL2",
		Description: "Tariff limit 2",
	},
	RegisterTariffLimit3: {
		Name:        "TL3",
		Description: "Tariff limit 3",
	},
	RegisterVolumeRegister1: {
		Name:        "V1",
		Description: "Volume register V1",
	},
	RegisterVolumeRegister2: {
		Name:        "V1",
		Description: "Volume register V2",
	},
	RegisterMassRegister1: {
		Name:        "M1",
		Description: "Mass register V1",
	},
	RegisterMassRegister2: {
		Name:        "M2",
		Description: "Mass register V2",
	},
	RegisterCurrentInFlow: {
		Name:        "FLOW1",
		Description: "Current flow in flow",
	},
	RegisterCurrentReturnFlow: {
		Name:        "FLOW2",
		Description: "Current flow in return flow",
	},
	RegisterCurrentPower: {
		Name:        "EFFEKT1",
		Description: "Current power calculated on the basis of V1-T1-T2",
	},
	RegisterInputVA: {
		Name:        "VA",
		Description: "Input register VA",
	},
	RegisterInputVB: {
		Name:        "VB",
		Description: "Input register VB",
	},
	RegisterCurrentInTemp: {
		Name:        "T1",
		Description: "Current flow temperature",
	},
	RegisterCurrentReturnTemp: {
		Name:        "T2",
		Description: "Current return flow temperature",
	},
	RegisterCurrentT3Temp: {
		Name:        "T3",
		Description: "Current temperature T3",
	},
	RegisterCurrentTempDifference: {
		Name:        "T1-T2",
		Description: "Current temperature difference",
	},
	RegisterCurrentInPressure: {
		Name:        "P1",
		Description: "Pressure in flow",
	},
	RegisterCurrentReturnPressure: {
		Name:        "P2",
		Description: "Pressure in return flow",
	},
	RegisterControlEnergy: {
		Name:        "E2",
		Description: "Control energy",
	},
	RegisterTapEnergy: {
		Name:        "E6",
		Description: "Energy register 6: Tap water energy",
	},
	RegisterYEnergy: {
		Name:        "E7",
		Description: "Energy register 7: Heat energy Y",
	},
	RegisterEnergy8: {
		Name:        "E8",
		Description: "Energy register 8: m³ x T1",
	},
	RegisterInfocode: {
		Name:        "INFO",
		Description: "Infocode, current",
	},
	RegisterMeterNoVB: {
		Name:        "METER NO VB",
		Description: "Meter no for VB",
	},
	RegisterEnergy9: {
		Name:        "E9",
		Description: "Energy register 9: m³ x T2",
	},
	RegisterMeterNo2: {
		Name:        "METER NO 2",
		Description: "Customer number (8 most important digits)",
	},
	RegisterInfoEvent: {
		Name:        "INFOEVENT",
		Description: "Info-event counter",
	},
	RegisterMeterNoVA: {
		Name:        "METER NO VA",
		Description: "Meter no for VA",
	},
	RegisterCurrentT4Temp: {
		Name:        "T4",
		Description: "Current temperature T4",
	},
	RegisterMaxFlowYTDDate: {
		Name:        "MAX FLOW1DATE/ÅR",
		Description: "Date for max flow this year",
	},
	RegisterMaxFlowYTDValue: {
		Name:        "MAX FLOW1/ÅR",
		Description: "Value for max flow this year",
	},
	RegisterMinFlowYTDDate: {
		Name:        "MIN FLOW1DATE/ÅR",
		Description: "Date for min flow this year",
	},
	RegisterMinFlowYTDValue: {
		Name:        "MIN FLOW1/ÅR",
		Description: "Value for min flow this year",
	},
	RegisterMaxPowerYTDDate: {
		Name:        "MAX EFFEKT1/ÅR",
		Description: "Date for max effekt this year",
	},
	RegisterMaxPowerYTDValue: {
		Name:        "MAX EFFEKT1/ÅR",
		Description: "Value for max effekt this year",
	},
	RegisterMinPowerYTDDate: {
		Name:        "MIN EFFEKT1DATE/ÅR",
		Description: "Date for min effekt this year",
	},
	RegisterMinPowerYTDValue: {
		Name:        "MIN EFFEKT1/ÅR",
		Description: "Value for min effekt this year",
	},
	RegisterMaxFlowMTDDate: {
		Name:        "MAX FLOW1DATE/MÅNED",
		Description: "Date for max flow this month",
	},
	RegisterMaxFlowMTDValue: {
		Name:        "MAX FLOW1/MÅNED",
		Description: "Value for max flow this month",
	},
	RegisterMinFlowMTDDate: {
		Name:        "MIN FLOW1DATE/MÅNED",
		Description: "Date for min flow this month",
	},
	RegisterMinFlowMTDValue: {
		Name:        "MIN FLOW1/MÅNED",
		Description: "Value for min flow this month",
	},
	RegisterMaxPowerMTDDate: {
		Name:        "MAX EFFEKT1/MÅNED",
		Description: "Date for max effekt this month",
	},
	RegisterMaxPowerMTDValue: {
		Name:        "MAX EFFEKT1/MÅNED",
		Description: "Value for max effekt this month",
	},
	RegisterMinPowerMTDDate: {
		Name:        "MIN EFFEKT1DATE/MÅNED",
		Description: "Date for min effekt this month",
	},
	RegisterMinPowerMTDValue: {
		Name:        "MIN EFFEKT1/MÅNED",
		Description: "Value for min effekt this month",
	},
	RegisterAvgT1YTD: {
		Name:        "AVR T1/ÅR",
		Description: "Year-to-date average for T1",
	},
	RegisterAvgT2YTD: {
		Name:        "AVR T2/ÅR",
		Description: "Year-to-date average for T2",
	},
	RegisterAvgT1MTD: {
		Name:        "AVR T1/MÅNED",
		Description: "Month-to-date average for T1",
	},
	RegisterAvgT2MTD: {
		Name:        "AVR T2/MÅNED",
		Description: "Month-to-date average for T2",
	},
	RegisterConfigNo1: {
		Name:        "CONFIG NO 1",
		Description: "Config no. DDDEE",
	},
	RegisterChecksum1: {
		Name:        "CHECK SUM 1",
		Description: "Software check sum",
	},
	RegisterConfigNo2: {
		Name:        "CONFIG NO 2",
		Description: "Config no. FFGGMN",
	},
	RegisterErrorHourCounter: {
		Name:        "ERRORHOURCOUNTER",
		Description: "Error hour counter",
	},
	RegisterInALiterImp: {
		Name:        "INA LITERIMP",
		Description: "Liter/imp for input A",
	},
	RegisterInBLiterImp: {
		Name:        "INB LITERIMP",
		Description: "Liter/imp for input B",
	},
	RegisterSerialNo: {
		Name:        "SERIE NO",
		Description: "Serial no. (unique number for each meter)",
	},
	RegisterClock: {
		Name:        "CLOCK",
		Description: "Current time (hhmmss)",
	},
	RegisterDate: {
		Name:        "DATE",
		Description: "Current date (YYMMDD)",
	},
	RegisterOperationalHours: {
		Name:        "HR",
		Description: "Operational hour counter",
	},
	RegisterSoftwareEdition: {
		Name:        "METER TYPE",
		Description: "Software edition",
	},
	RegisterMeterNo1: {
		Name:        "METER NO 1",
		Description: "Customer number (LS)",
	},
}

func (r RegisterID) String() string {
	if ri, ok := regInfo[r]; ok {
		return fmt.Sprintf("%s (%s)", ri.Name, ri.Description)
	}

	return "Unknown"
}
