/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.9 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// PatchedWritablePowerOutletTemplateRequestType * `iec-60320-c5` - C5 * `iec-60320-c7` - C7 * `iec-60320-c13` - C13 * `iec-60320-c15` - C15 * `iec-60320-c19` - C19 * `iec-60320-c21` - C21 * `iec-60309-p-n-e-4h` - P+N+E 4H * `iec-60309-p-n-e-6h` - P+N+E 6H * `iec-60309-p-n-e-9h` - P+N+E 9H * `iec-60309-2p-e-4h` - 2P+E 4H * `iec-60309-2p-e-6h` - 2P+E 6H * `iec-60309-2p-e-9h` - 2P+E 9H * `iec-60309-3p-e-4h` - 3P+E 4H * `iec-60309-3p-e-6h` - 3P+E 6H * `iec-60309-3p-e-9h` - 3P+E 9H * `iec-60309-3p-n-e-4h` - 3P+N+E 4H * `iec-60309-3p-n-e-6h` - 3P+N+E 6H * `iec-60309-3p-n-e-9h` - 3P+N+E 9H * `iec-60906-1` - IEC 60906-1 * `nbr-14136-10a` - 2P+T 10A (NBR 14136) * `nbr-14136-20a` - 2P+T 20A (NBR 14136) * `nema-1-15r` - NEMA 1-15R * `nema-5-15r` - NEMA 5-15R * `nema-5-20r` - NEMA 5-20R * `nema-5-30r` - NEMA 5-30R * `nema-5-50r` - NEMA 5-50R * `nema-6-15r` - NEMA 6-15R * `nema-6-20r` - NEMA 6-20R * `nema-6-30r` - NEMA 6-30R * `nema-6-50r` - NEMA 6-50R * `nema-10-30r` - NEMA 10-30R * `nema-10-50r` - NEMA 10-50R * `nema-14-20r` - NEMA 14-20R * `nema-14-30r` - NEMA 14-30R * `nema-14-50r` - NEMA 14-50R * `nema-14-60r` - NEMA 14-60R * `nema-15-15r` - NEMA 15-15R * `nema-15-20r` - NEMA 15-20R * `nema-15-30r` - NEMA 15-30R * `nema-15-50r` - NEMA 15-50R * `nema-15-60r` - NEMA 15-60R * `nema-l1-15r` - NEMA L1-15R * `nema-l5-15r` - NEMA L5-15R * `nema-l5-20r` - NEMA L5-20R * `nema-l5-30r` - NEMA L5-30R * `nema-l5-50r` - NEMA L5-50R * `nema-l6-15r` - NEMA L6-15R * `nema-l6-20r` - NEMA L6-20R * `nema-l6-30r` - NEMA L6-30R * `nema-l6-50r` - NEMA L6-50R * `nema-l10-30r` - NEMA L10-30R * `nema-l14-20r` - NEMA L14-20R * `nema-l14-30r` - NEMA L14-30R * `nema-l14-50r` - NEMA L14-50R * `nema-l14-60r` - NEMA L14-60R * `nema-l15-20r` - NEMA L15-20R * `nema-l15-30r` - NEMA L15-30R * `nema-l15-50r` - NEMA L15-50R * `nema-l15-60r` - NEMA L15-60R * `nema-l21-20r` - NEMA L21-20R * `nema-l21-30r` - NEMA L21-30R * `nema-l22-30r` - NEMA L22-30R * `CS6360C` - CS6360C * `CS6364C` - CS6364C * `CS8164C` - CS8164C * `CS8264C` - CS8264C * `CS8364C` - CS8364C * `CS8464C` - CS8464C * `ita-e` - ITA Type E (CEE 7/5) * `ita-f` - ITA Type F (CEE 7/3) * `ita-g` - ITA Type G (BS 1363) * `ita-h` - ITA Type H * `ita-i` - ITA Type I * `ita-j` - ITA Type J * `ita-k` - ITA Type K * `ita-l` - ITA Type L (CEI 23-50) * `ita-m` - ITA Type M (BS 546) * `ita-n` - ITA Type N * `ita-o` - ITA Type O * `ita-multistandard` - ITA Multistandard * `usb-a` - USB Type A * `usb-micro-b` - USB Micro B * `usb-c` - USB Type C * `molex-micro-fit-1x2` - Molex Micro-Fit 1x2 * `molex-micro-fit-2x2` - Molex Micro-Fit 2x2 * `molex-micro-fit-2x4` - Molex Micro-Fit 2x4 * `dc-terminal` - DC Terminal * `hdot-cx` - HDOT Cx * `saf-d-grid` - Saf-D-Grid * `neutrik-powercon-20a` - Neutrik powerCON (20A) * `neutrik-powercon-32a` - Neutrik powerCON (32A) * `neutrik-powercon-true1` - Neutrik powerCON TRUE1 * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP * `ubiquiti-smartpower` - Ubiquiti SmartPower * `hardwired` - Hardwired * `other` - Other
type PatchedWritablePowerOutletTemplateRequestType string

// List of PatchedWritablePowerOutletTemplateRequest_type
const (
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60320_C5               PatchedWritablePowerOutletTemplateRequestType = "iec-60320-c5"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60320_C7               PatchedWritablePowerOutletTemplateRequestType = "iec-60320-c7"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60320_C13              PatchedWritablePowerOutletTemplateRequestType = "iec-60320-c13"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60320_C15              PatchedWritablePowerOutletTemplateRequestType = "iec-60320-c15"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60320_C19              PatchedWritablePowerOutletTemplateRequestType = "iec-60320-c19"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60320_C21              PatchedWritablePowerOutletTemplateRequestType = "iec-60320-c21"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_P_N_E_4H         PatchedWritablePowerOutletTemplateRequestType = "iec-60309-p-n-e-4h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_P_N_E_6H         PatchedWritablePowerOutletTemplateRequestType = "iec-60309-p-n-e-6h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_P_N_E_9H         PatchedWritablePowerOutletTemplateRequestType = "iec-60309-p-n-e-9h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_2P_E_4H          PatchedWritablePowerOutletTemplateRequestType = "iec-60309-2p-e-4h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_2P_E_6H          PatchedWritablePowerOutletTemplateRequestType = "iec-60309-2p-e-6h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_2P_E_9H          PatchedWritablePowerOutletTemplateRequestType = "iec-60309-2p-e-9h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_3P_E_4H          PatchedWritablePowerOutletTemplateRequestType = "iec-60309-3p-e-4h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_3P_E_6H          PatchedWritablePowerOutletTemplateRequestType = "iec-60309-3p-e-6h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_3P_E_9H          PatchedWritablePowerOutletTemplateRequestType = "iec-60309-3p-e-9h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_3P_N_E_4H        PatchedWritablePowerOutletTemplateRequestType = "iec-60309-3p-n-e-4h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_3P_N_E_6H        PatchedWritablePowerOutletTemplateRequestType = "iec-60309-3p-n-e-6h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60309_3P_N_E_9H        PatchedWritablePowerOutletTemplateRequestType = "iec-60309-3p-n-e-9h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_IEC_60906_1                PatchedWritablePowerOutletTemplateRequestType = "iec-60906-1"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NBR_14136_10A              PatchedWritablePowerOutletTemplateRequestType = "nbr-14136-10a"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NBR_14136_20A              PatchedWritablePowerOutletTemplateRequestType = "nbr-14136-20a"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_1_15R                 PatchedWritablePowerOutletTemplateRequestType = "nema-1-15r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_5_15R                 PatchedWritablePowerOutletTemplateRequestType = "nema-5-15r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_5_20R                 PatchedWritablePowerOutletTemplateRequestType = "nema-5-20r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_5_30R                 PatchedWritablePowerOutletTemplateRequestType = "nema-5-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_5_50R                 PatchedWritablePowerOutletTemplateRequestType = "nema-5-50r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_6_15R                 PatchedWritablePowerOutletTemplateRequestType = "nema-6-15r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_6_20R                 PatchedWritablePowerOutletTemplateRequestType = "nema-6-20r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_6_30R                 PatchedWritablePowerOutletTemplateRequestType = "nema-6-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_6_50R                 PatchedWritablePowerOutletTemplateRequestType = "nema-6-50r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_10_30R                PatchedWritablePowerOutletTemplateRequestType = "nema-10-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_10_50R                PatchedWritablePowerOutletTemplateRequestType = "nema-10-50r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_14_20R                PatchedWritablePowerOutletTemplateRequestType = "nema-14-20r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_14_30R                PatchedWritablePowerOutletTemplateRequestType = "nema-14-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_14_50R                PatchedWritablePowerOutletTemplateRequestType = "nema-14-50r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_14_60R                PatchedWritablePowerOutletTemplateRequestType = "nema-14-60r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_15_15R                PatchedWritablePowerOutletTemplateRequestType = "nema-15-15r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_15_20R                PatchedWritablePowerOutletTemplateRequestType = "nema-15-20r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_15_30R                PatchedWritablePowerOutletTemplateRequestType = "nema-15-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_15_50R                PatchedWritablePowerOutletTemplateRequestType = "nema-15-50r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_15_60R                PatchedWritablePowerOutletTemplateRequestType = "nema-15-60r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L1_15R                PatchedWritablePowerOutletTemplateRequestType = "nema-l1-15r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L5_15R                PatchedWritablePowerOutletTemplateRequestType = "nema-l5-15r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L5_20R                PatchedWritablePowerOutletTemplateRequestType = "nema-l5-20r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L5_30R                PatchedWritablePowerOutletTemplateRequestType = "nema-l5-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L5_50R                PatchedWritablePowerOutletTemplateRequestType = "nema-l5-50r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L6_15R                PatchedWritablePowerOutletTemplateRequestType = "nema-l6-15r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L6_20R                PatchedWritablePowerOutletTemplateRequestType = "nema-l6-20r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L6_30R                PatchedWritablePowerOutletTemplateRequestType = "nema-l6-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L6_50R                PatchedWritablePowerOutletTemplateRequestType = "nema-l6-50r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L10_30R               PatchedWritablePowerOutletTemplateRequestType = "nema-l10-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L14_20R               PatchedWritablePowerOutletTemplateRequestType = "nema-l14-20r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L14_30R               PatchedWritablePowerOutletTemplateRequestType = "nema-l14-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L14_50R               PatchedWritablePowerOutletTemplateRequestType = "nema-l14-50r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L14_60R               PatchedWritablePowerOutletTemplateRequestType = "nema-l14-60r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L15_20R               PatchedWritablePowerOutletTemplateRequestType = "nema-l15-20r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L15_30R               PatchedWritablePowerOutletTemplateRequestType = "nema-l15-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L15_50R               PatchedWritablePowerOutletTemplateRequestType = "nema-l15-50r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L15_60R               PatchedWritablePowerOutletTemplateRequestType = "nema-l15-60r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L21_20R               PatchedWritablePowerOutletTemplateRequestType = "nema-l21-20r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L21_30R               PatchedWritablePowerOutletTemplateRequestType = "nema-l21-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEMA_L22_30R               PatchedWritablePowerOutletTemplateRequestType = "nema-l22-30r"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_CS6360_C                   PatchedWritablePowerOutletTemplateRequestType = "CS6360C"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_CS6364_C                   PatchedWritablePowerOutletTemplateRequestType = "CS6364C"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_CS8164_C                   PatchedWritablePowerOutletTemplateRequestType = "CS8164C"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_CS8264_C                   PatchedWritablePowerOutletTemplateRequestType = "CS8264C"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_CS8364_C                   PatchedWritablePowerOutletTemplateRequestType = "CS8364C"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_CS8464_C                   PatchedWritablePowerOutletTemplateRequestType = "CS8464C"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_E                      PatchedWritablePowerOutletTemplateRequestType = "ita-e"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_F                      PatchedWritablePowerOutletTemplateRequestType = "ita-f"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_G                      PatchedWritablePowerOutletTemplateRequestType = "ita-g"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_H                      PatchedWritablePowerOutletTemplateRequestType = "ita-h"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_I                      PatchedWritablePowerOutletTemplateRequestType = "ita-i"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_J                      PatchedWritablePowerOutletTemplateRequestType = "ita-j"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_K                      PatchedWritablePowerOutletTemplateRequestType = "ita-k"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_L                      PatchedWritablePowerOutletTemplateRequestType = "ita-l"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_M                      PatchedWritablePowerOutletTemplateRequestType = "ita-m"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_N                      PatchedWritablePowerOutletTemplateRequestType = "ita-n"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_O                      PatchedWritablePowerOutletTemplateRequestType = "ita-o"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_ITA_MULTISTANDARD          PatchedWritablePowerOutletTemplateRequestType = "ita-multistandard"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_USB_A                      PatchedWritablePowerOutletTemplateRequestType = "usb-a"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_USB_MICRO_B                PatchedWritablePowerOutletTemplateRequestType = "usb-micro-b"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_USB_C                      PatchedWritablePowerOutletTemplateRequestType = "usb-c"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_MOLEX_MICRO_FIT_1X2        PatchedWritablePowerOutletTemplateRequestType = "molex-micro-fit-1x2"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_MOLEX_MICRO_FIT_2X2        PatchedWritablePowerOutletTemplateRequestType = "molex-micro-fit-2x2"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_MOLEX_MICRO_FIT_2X4        PatchedWritablePowerOutletTemplateRequestType = "molex-micro-fit-2x4"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_DC_TERMINAL                PatchedWritablePowerOutletTemplateRequestType = "dc-terminal"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_HDOT_CX                    PatchedWritablePowerOutletTemplateRequestType = "hdot-cx"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_SAF_D_GRID                 PatchedWritablePowerOutletTemplateRequestType = "saf-d-grid"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEUTRIK_POWERCON_20A       PatchedWritablePowerOutletTemplateRequestType = "neutrik-powercon-20a"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEUTRIK_POWERCON_32A       PatchedWritablePowerOutletTemplateRequestType = "neutrik-powercon-32a"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEUTRIK_POWERCON_TRUE1     PatchedWritablePowerOutletTemplateRequestType = "neutrik-powercon-true1"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_NEUTRIK_POWERCON_TRUE1_TOP PatchedWritablePowerOutletTemplateRequestType = "neutrik-powercon-true1-top"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_UBIQUITI_SMARTPOWER        PatchedWritablePowerOutletTemplateRequestType = "ubiquiti-smartpower"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_HARDWIRED                  PatchedWritablePowerOutletTemplateRequestType = "hardwired"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_OTHER                      PatchedWritablePowerOutletTemplateRequestType = "other"
	PATCHEDWRITABLEPOWEROUTLETTEMPLATEREQUESTTYPE_EMPTY                      PatchedWritablePowerOutletTemplateRequestType = ""
)

// All allowed values of PatchedWritablePowerOutletTemplateRequestType enum
var AllowedPatchedWritablePowerOutletTemplateRequestTypeEnumValues = []PatchedWritablePowerOutletTemplateRequestType{
	"iec-60320-c5",
	"iec-60320-c7",
	"iec-60320-c13",
	"iec-60320-c15",
	"iec-60320-c19",
	"iec-60320-c21",
	"iec-60309-p-n-e-4h",
	"iec-60309-p-n-e-6h",
	"iec-60309-p-n-e-9h",
	"iec-60309-2p-e-4h",
	"iec-60309-2p-e-6h",
	"iec-60309-2p-e-9h",
	"iec-60309-3p-e-4h",
	"iec-60309-3p-e-6h",
	"iec-60309-3p-e-9h",
	"iec-60309-3p-n-e-4h",
	"iec-60309-3p-n-e-6h",
	"iec-60309-3p-n-e-9h",
	"iec-60906-1",
	"nbr-14136-10a",
	"nbr-14136-20a",
	"nema-1-15r",
	"nema-5-15r",
	"nema-5-20r",
	"nema-5-30r",
	"nema-5-50r",
	"nema-6-15r",
	"nema-6-20r",
	"nema-6-30r",
	"nema-6-50r",
	"nema-10-30r",
	"nema-10-50r",
	"nema-14-20r",
	"nema-14-30r",
	"nema-14-50r",
	"nema-14-60r",
	"nema-15-15r",
	"nema-15-20r",
	"nema-15-30r",
	"nema-15-50r",
	"nema-15-60r",
	"nema-l1-15r",
	"nema-l5-15r",
	"nema-l5-20r",
	"nema-l5-30r",
	"nema-l5-50r",
	"nema-l6-15r",
	"nema-l6-20r",
	"nema-l6-30r",
	"nema-l6-50r",
	"nema-l10-30r",
	"nema-l14-20r",
	"nema-l14-30r",
	"nema-l14-50r",
	"nema-l14-60r",
	"nema-l15-20r",
	"nema-l15-30r",
	"nema-l15-50r",
	"nema-l15-60r",
	"nema-l21-20r",
	"nema-l21-30r",
	"nema-l22-30r",
	"CS6360C",
	"CS6364C",
	"CS8164C",
	"CS8264C",
	"CS8364C",
	"CS8464C",
	"ita-e",
	"ita-f",
	"ita-g",
	"ita-h",
	"ita-i",
	"ita-j",
	"ita-k",
	"ita-l",
	"ita-m",
	"ita-n",
	"ita-o",
	"ita-multistandard",
	"usb-a",
	"usb-micro-b",
	"usb-c",
	"molex-micro-fit-1x2",
	"molex-micro-fit-2x2",
	"molex-micro-fit-2x4",
	"dc-terminal",
	"hdot-cx",
	"saf-d-grid",
	"neutrik-powercon-20a",
	"neutrik-powercon-32a",
	"neutrik-powercon-true1",
	"neutrik-powercon-true1-top",
	"ubiquiti-smartpower",
	"hardwired",
	"other",
	"",
}

func (v *PatchedWritablePowerOutletTemplateRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritablePowerOutletTemplateRequestType(value)
	for _, existing := range AllowedPatchedWritablePowerOutletTemplateRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritablePowerOutletTemplateRequestType", value)
}

// NewPatchedWritablePowerOutletTemplateRequestTypeFromValue returns a pointer to a valid PatchedWritablePowerOutletTemplateRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritablePowerOutletTemplateRequestTypeFromValue(v string) (*PatchedWritablePowerOutletTemplateRequestType, error) {
	ev := PatchedWritablePowerOutletTemplateRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritablePowerOutletTemplateRequestType: valid values are %v", v, AllowedPatchedWritablePowerOutletTemplateRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritablePowerOutletTemplateRequestType) IsValid() bool {
	for _, existing := range AllowedPatchedWritablePowerOutletTemplateRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritablePowerOutletTemplateRequest_type value
func (v PatchedWritablePowerOutletTemplateRequestType) Ptr() *PatchedWritablePowerOutletTemplateRequestType {
	return &v
}

type NullablePatchedWritablePowerOutletTemplateRequestType struct {
	value *PatchedWritablePowerOutletTemplateRequestType
	isSet bool
}

func (v NullablePatchedWritablePowerOutletTemplateRequestType) Get() *PatchedWritablePowerOutletTemplateRequestType {
	return v.value
}

func (v *NullablePatchedWritablePowerOutletTemplateRequestType) Set(val *PatchedWritablePowerOutletTemplateRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerOutletTemplateRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerOutletTemplateRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerOutletTemplateRequestType(val *PatchedWritablePowerOutletTemplateRequestType) *NullablePatchedWritablePowerOutletTemplateRequestType {
	return &NullablePatchedWritablePowerOutletTemplateRequestType{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerOutletTemplateRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerOutletTemplateRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
