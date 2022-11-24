package conn

import (
	"errors"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/accounting"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/act"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/cdr"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/dui"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/gdt"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/inv"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/inventory"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/kendaql"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/mes"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/per"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/plm"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/prd"
	pp "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/productionplanning"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/sas"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/sec"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/utility"
)

// DM implements backend services with DM.
type DM struct {
	client *Client
}

// NewDM returns a new instance of DM services.
func NewDM() *DM {
	return &DM{}
}

// Name returns the name of the services implementation.
func (s *DM) Name() string {
	return "DM"
}

// Connect establishes backend connections.
func (s *DM) Connect(opts *Options) error {
	if s.client != nil {
		return errors.New("conn error:DM client already initialized")
	}

	client, err := NewClient(opts)
	if err != nil {
		return err
	}
	s.client = client

	return nil
}

// Close closes client connections to all backends.
func (s *DM) Close() error {
	return s.client.Close()
}

// Act returns ACT service client.
func (s *DM) Act() act.ACTClient {
	return act.NewACTClient(s.client.conn)
}

// Accounting returns Accounting service client.
func (s *DM) Accounting() accounting.AccountingClient {
	return accounting.NewAccountingClient(s.client.conn)
}

// Cdr returns CDR service client.
func (s *DM) Cdr() cdr.CDRClient {
	return cdr.NewCDRClient(s.client.conn)
}

// Gdt returns GDT service client.
func (s *DM) Gdt() gdt.GDTClient {
	return gdt.NewGDTClient(s.client.conn)
}

// Inv returns INV service client.
func (s *DM) Inv() inv.INVClient {
	return inv.NewINVClient(s.client.conn)
}

// Inventory returns Inventory service client.
func (s *DM) Inventory() inventory.InventoryClient {
	return inventory.NewInventoryClient(s.client.conn)
}

// Mes returns MES service client.
func (s *DM) Mes() mes.MESClient {
	return mes.NewMESClient(s.client.conn)
}

// Per returns PER service client.
func (s *DM) Per() per.PERClient {
	return per.NewPERClient(s.client.conn)
}

// Plm returns PLM service client.
func (s *DM) Plm() plm.PLMClient {
	return plm.NewPLMClient(s.client.conn)
}

// Prd returns PRD service client.
func (s *DM) Prd() prd.PRDClient {
	return prd.NewPRDClient(s.client.conn)
}

// ProductionPlanning ProductionPlanning service client.
func (s *DM) ProductionPlanning() pp.ProductionPlanningClient {
	return pp.NewProductionPlanningClient(s.client.conn)
}

// Rs returns RS service client.
func (s *DM) Rs() rs.RSClient {
	return rs.NewRSClient(s.client.conn)
}

// Sas returns SAS service client.
func (s *DM) Sas() sas.SASClient {
	return sas.NewSASClient(s.client.conn)
}

// Sec returns SEC service client.
func (s *DM) Sec() sec.SECClient {
	return sec.NewSECClient(s.client.conn)
}

// Utility returns UTILITY service client.
func (s *DM) Utility() utility.UtilityClient {
	return utility.NewUtilityClient(s.client.conn)
}

// KendaQL returns KENDAQL service client.
func (s *DM) KendaQL() kendaql.KendaQLClient {
	return kendaql.NewKendaQLClient(s.client.conn)
}

// Auth returns Auth service client.
func (s *DM) Auth() utility.AuthenticationClient {
	return utility.NewAuthenticationClient(s.client.conn)
}

// Dui returns DUI service client.
func (s *DM) Dui() dui.DUIClient {
	return dui.NewDUIClient(s.client.conn)
}
