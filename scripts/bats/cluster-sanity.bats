#!/usr/bin/env bats

@test "Check Replicas Count" {
  count="$(docker ps | grep -c sort)"
  [ "$count" -eq 3 ]
}

@test "Check Replicas Are Available" {
  ports="$(docker ps | awk '/sort/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	for port in "${ports_list[@]}"; do
			response="$(curl -sS -X GET http://$port/)"
			[ "$response" == "Hello World Sort Node" ]
	done 
}
