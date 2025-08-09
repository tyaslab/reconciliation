Reconciliation
==============

Reconciliation makes easy. Example csv file available at `fixture` directory.

Run Server
==========

    $ go run main.go

Access `POST http://127.0.0.1:5000/reconcile`

Example request:

```
curl --location 'http://127.0.0.1:5000/reconcile' \
--form 'date_from="2025-08-01"' \
--form 'date_to="2025-08-01"' \
--form 'transaction_file=@"/Users/adityadarmawan/Projects/reconciliation/fixture/matched/transaction.csv"' \
--form 'bank_statement_file=@"/Users/adityadarmawan/Projects/reconciliation/fixture/matched/bank.csv"'
```

Example response:
```
{
    "processed": 8,
    "matched": 4,
    "system_unmatched": 0,
    "bank_unmatched": 0,
    "discrepancies": 0
}
```

Test
====

    $ go test adityadarmawan.id/recon/app -v

