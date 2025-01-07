# ----------------------------------------
# Makefile for managing multiple services
# ----------------------------------------

# ====== Variables ======
FRONTEND_DIR=./ts/web/packages/frontend
BFF_DIR=./ts/web/packages/bff
ACCOUNT_MGR_DIR=./go/app/account/cmd/account-mgr
TICKET_MGR_DIR=./go/app/ticket/cmd/ticket-mgr
TICKET_QUERIER_DIR=./go/app/ticket/cmd/ticket-querier

# ====== Commands ======
FRONTEND_CMD=bun run --cwd $(FRONTEND_DIR) dev
BFF_CMD=bun run --cwd $(BFF_DIR) dev
ACCOUNT_MGR_CMD=cd $(ACCOUNT_MGR_DIR) && go run main.go
TICKET_MGR_CMD=cd $(TICKET_MGR_DIR) && go run main.go
TICKET_QUERIER_CMD=cd $(TICKET_QUERIER_DIR) && go run main.go

# ====== Start Services ======
.PHONY: frontend bff account-mgr ticket-mgr ticket-querier start stop

frontend:
	@echo "🚀 Starting Frontend Service..."
	$(FRONTEND_CMD)

bff:
	@echo "🚀 Starting BFF Service..."
	$(BFF_CMD)

account-mgr:
	@echo "🚀 Starting Account Manager Service..."
	$(ACCOUNT_MGR_CMD)

ticket-mgr:
	@echo "🚀 Starting Ticket Manager Service..."
	$(TICKET_MGR_CMD)

ticket-querier:
	@echo "🚀 Starting Ticket Querier Service..."
	$(TICKET_QUERIER_CMD)

# ====== Start All Services in Parallel ======
start:
	@echo "🚀 Starting All Services..."
	@make -j 5 frontend bff account-mgr ticket-mgr ticket-querier
