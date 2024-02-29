#!/bin/bash

# Define SSH connection details


# Use sshpass for the SSH connection to run commands and retrieve the cookie
# COOKIE=$(sshpass -p "$PASSWORD" ssh -o StrictHostKeyChecking=no "$USERNAME@$HOST" '
COOKIE=$(sshpass -p "hessam123" ssh -o StrictHostKeyChecking=no "hessam@10.10.5.103" '
cd nap
./script-auth.sh
cat cookie.curl
')

# Parse the TGC and JSESID values from the cookie
TGC_VALUE=$(echo "${COOKIE}" | grep -oP 'TGC\s+\K.+')
JSESID_VALUE=$(echo "${COOKIE}" | grep -oP 'JSESID85d25c\s+\K.+')


# Update the cookie.curl file with the new TGC and JSESID values
{
    echo "10.79.23.42   FALSE   /cas    TRUE    0       TGC     ${TGC_VALUE}"
    echo "10.79.23.42   FALSE   /onc    FALSE   0       JSESID85d25c    ${JSESID_VALUE}"
} > cookie.curl

echo "Updated cookie.curl with new TGC and JSESID values."
# Output the extracted keys
echo "$TGC_VALUE"
echo "$JSESID_VALUE"
