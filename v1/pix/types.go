package pix

import v1 "github.com/AbacatePay/go-types/v1"

type PIX = v1.APIQRCodePIX

type CreateQrCodePIXBody = v1.RESTPostCreateQRCodePixBody

type CreateQrCodePIXData = v1.RESTPostCreateQRCodePixData

type CheckPIXStatusData = v1.RESTGetCheckQRCodePixStatusData

type SimulatePaymentData = v1.RESTPostSimulatePaymentData
