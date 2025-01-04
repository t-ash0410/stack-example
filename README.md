# stack-example

## gRPC call example

Create ticket

```zsh
buf curl --http2-prior-knowledge --protocol grpc \
  http://localhost:8080/ticketmgr.v1.TicketMgrService/CreateTicket \
  -d '{
    "requested_by":"8ea79f88-5b4b-4df6-b438-81a2ccf6b09f",
    "title": "Some Ticket",
    "description": "Some ticket description.",
    "deadline": "2024-12-31T15:00:00+00:00"
  }'
```

Update ticket

```zsh
buf curl --http2-prior-knowledge --protocol grpc \
  http://localhost:8080/ticketmgr.v1.TicketMgrService/UpdateTicket \
  -d '{
    "ticket_id": "some ticket id",
    "requested_by":"4e770fc1-0977-4ea9-911a-d67d4185817e",
    "title": "Updated Ticket",
    "description": "Updated ticket description.",
    "deadline": "2025-01-19T15:00:00+00:00"
  }'
```

Delete ticket

```zsh
buf curl --http2-prior-knowledge --protocol grpc \
  http://localhost:8080/ticketmgr.v1.TicketMgrService/DeleteTicket \
  -d '{
    "ticket_id": "some ticket id"
  }'
```

List tickets

```zsh
buf curl --http2-prior-knowledge --protocol grpc \
  http://localhost:8081/ticketquerier.v1.TicketQuerierService/QueryTickets \
  -d '{
    "requested_by": "8ea79f88-5b4b-4df6-b438-81a2ccf6b09f"
  }'
```

Get ticket

```zsh
buf curl --http2-prior-knowledge --protocol grpc \
  http://localhost:8081/ticketquerier.v1.TicketQuerierService/GetTicketById \
  -d '{
    "ticket_id": "some ticket id"
  }'
```
