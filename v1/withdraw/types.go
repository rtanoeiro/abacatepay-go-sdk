package withdraw

import v1 "github.com/AbacatePay/go-types/v1"

type CreateWithdrawBody = v1.RESTPostCreateNewWithdrawBody

type CreateWithdrawResponse = v1.RESTPostCreateNewChargeData

type GetWithdrawQuery = v1.RESTGetSearchWithdrawQueryParams

// TODO: Fix this later
type GetWithdrawResponse = v1.APIResponse[v1.APIWithdraw]

type ListWithdrawsData = v1.RESTGetListWithdrawsData
