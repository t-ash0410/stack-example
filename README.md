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
