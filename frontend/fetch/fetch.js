import fetch from "node-fetch"

const data = { 
	name: 'Pedro',
	age:24,
	residence: 'Boston',
	birthplace: 'New York',
	occupation: 'Archeologist'
};

fetch('http://localhost:8080/v1/investigator', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(data),
})
.then(response => response.json())
.then(data => {
  console.log('Success:', data);
})
.catch((error) => {
  console.error('Error:', error);
});

