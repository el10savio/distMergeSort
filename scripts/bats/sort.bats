#!/usr/bin/env bats

@test "Sort Empty List" {
	ports="$(docker ps | awk '/sort/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	node1="${ports_list[0]}"

	response="$(curl -sS -XPOST http://$node1/sort --data '{"values": []}')" && [ "$response" == "[]" ]
}

@test "Sort Basic List" {
	ports="$(docker ps | awk '/sort/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	node1="${ports_list[0]}"

	response="$(curl -sS -XPOST http://$node1/sort --data '{"values": [67,44,9,21,14,17,46,94,60,97]}')" && [ "$response" == "[9,14,17,21,44,46,60,67,94,97]" ]
}

@test "Sort Basic List II" {
	ports="$(docker ps | awk '/sort/ {print $1}' | xargs -I {} docker port {} 8080 | sed ':a;N;$!ba;s/\n/,/g' | sort)"
	IFS=', ' read -r -a ports_list <<< "$ports"

	node1="${ports_list[0]}"

	response="$(curl -sS -XPOST http://$node1/sort --data '{"values": [67,44,9,21,14,17,46,94,60]}')" && [ "$response" == "[9,14,17,21,44,46,60,67,94]" ]
}
