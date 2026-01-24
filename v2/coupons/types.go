package coupons

import v2 "github.com/AbacatePay/go-types/v2"

type Coupon = v2.APICoupon

type CreateCouponBody = v2.RESTPostCreateCouponBody

type CreateCouponData = v2.RESTPostCreateCouponData

type ListCouponsData = v2.RESTGetListCouponsData

type GetCouponData = v2.RESTGetCouponData

type ToggleStatusData = v2.RESTPatchToggleCouponStatusData

type DeleteCouponData = v2.RESTDeleteCouponData
