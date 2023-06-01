#!/bin/bash
source ../common.sh
echo "cluster-k3s: TCP & SCTP Multihoming combined"

if [ "$1" ]; then
  KUBECONFIG="$1"
fi

# Set space as the delimiter
IFS=' '

for((i=0; i<120; i++))
do
  extLB=$(sudo kubectl $KUBECONFIG get svc | grep "nginx-lb")
  read -a strarr <<< "$extLB"
  len=${#strarr[*]}
  if [[ $((len)) -lt 6 ]]; then
    echo "Can't find nginx-lb service"
    sleep 1
    continue
  fi 
  if [[ ${strarr[3]} != *"none"* ]]; then
    extIP=${strarr[3]}
    port=${strarr[4]}
    break
  fi
  echo "No external LB allocated"
  sleep 1
done

## Any routing updates  ??
sleep 30

echo "TCP service nginx-lb -> $extIP:$port"
out=$($hexec user curl -s --connect-timeout 10 http://$extIP:80) 

if [[ ${out} == *"Welcome to nginx"* ]]; then
  echo "cluster-k3s TCP service nginx-lb [OK]"
else
  echo "cluster-k3s TCP service nginx-lb [FAILED]"
  ## Dump some debug info
  echo "llb1 lb-info"
  $dexec llb1 loxicmd get lb
  echo "llb1 route-info"
  $dexec llb1 ip route
  echo "llb2 lb-info"
  $dexec llb2 loxicmd get lb
  echo "llb2 route-info"
  $dexec llb2 ip route
  echo "r1 route-info"
  $dexec r1 ip route
  exit 1
fi

out=$($hexec user curl -s --connect-timeout 10 http://$extIP:55002) 

if [[ ${out} == *"Welcome to nginx"* ]]; then
  echo "cluster-k3s TCP service nginx-lb (kube-loxilb) [OK]"
else
  echo "cluster-k3s TCP service nginx-lb (kube-loxilb) [FAILED]"
  ## Dump some debug info
  echo "llb1 lb-info"
  $dexec llb1 loxicmd get lb
  echo "llb1 route-info"
  $dexec llb1 ip route
  echo "llb2 lb-info"
  $dexec llb2 loxicmd get lb
  echo "llb2 route-info"
  $dexec llb2 ip route
  echo "r1 route-info"
  $dexec r1 ip route
  exit 1
fi

for((i=0; i<120; i++))
do
  extLB=$(sudo kubectl $KUBECONFIG get svc | grep "sctp-lb1")
  read -a strarr <<< "$extLB"
  len=${#strarr[*]}
  if [[ $((len)) -lt 6 ]]; then
    echo "Can't find sctp-lb1 service"
    sleep 1
    continue
  fi 
  if [[ ${strarr[3]} != *"none"* ]]; then
    extIP=${strarr[3]}
    port=${strarr[4]}
    break
  fi
  echo "No external LB allocated"
  sleep 1
done

echo "SCTP Multihoming service sctp-lb1 -> $extIP:$port"

$hexec user sctp_darn -H 1.1.1.1 -h 123.123.123.1 -p 55002 -s < input > output
sleep 5
exp="New connection, peer addresses
123.123.123.1:55002
124.124.124.1:55002
125.125.125.1:55002"

res=`cat output | grep -A 3 "New connection, peer addresses"`
if [[ "$res" == "$exp" ]]; then
    echo $res
    echo "cluster-k3s SCTP Multihoming service sctp-lb1 (kube-loxilb) [OK]"
else
    echo "cluster-k3s SCTP Multihoming service sctp-lb1 (kube-loxilb) [NOK]"
    echo "Expected : $exp"
    echo "Received : $res"
fi
sudo rm -rf output