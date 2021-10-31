#!/usr/bin/env bats

@test "Sort Empty List" {
  ports="$(docker ps | awk '/sort/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	node1="${ports_list[0]}"

	response="$(curl -sS -XPOST http://$node1/sort --data '{"values": []}')" && [ "$response" == "[]" ]
}
