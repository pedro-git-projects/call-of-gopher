import fetch from 'node-fetch';

const body = {
	name: 'Pedro',
	age:24,
	residence: 'Boston',
	birthplace: 'New York',
	occupation: 'Archeologist'
};

const response = await fetch('http://localhost:8080/v1/investigator', {
	method: 'post',
	body: JSON.stringify(body),
	headers: {'Content-Type': 'application/json'}
});
const data = await response.json();

console.log(data)
console.log(data.investigator.name)
console.log(data.investigator.str)

