#!/bin/bash

sshpass -pibof ssh -o StrictHostKeyChecking=no root@$1 "bash -s" < ./run_ibofos.sh

exit 0