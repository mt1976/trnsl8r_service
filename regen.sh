#!/bin/bash
# file: run.sh
echo "Regenerating DAO code..."
echo "-----------------------"
echo ""
echo "Generating TextStore DAO..."
go run github.com/mt1976/frantic-amphora/cmd/dao-gen -out ./app/dao/textStore -pkg textStore -type TextStore -table TextStore -namespace main -uri text -force
echo ""
echo "DAO code generation complete."
echo "-----------------------"
echo "Regeneration complete."