require('dotenv').config();
const axios = require('axios');
const dayjs = require('dayjs');

const endpoint = process.env.ENDPOINT;
const uuid = process.env.UUID;

if (!endpoint || !uuid) {
  console.error('Missing required environment variables:');
  if (!endpoint) console.error(' - ENDPOINT is not defined');
  if (!uuid) console.error(' - UUID is not defined');
  process.exit(1);
}

const timestamp = dayjs().format('YYYY-MM-DD HH:mm:ss');

const payload = {
  uuid: uuid,
  timestamp: timestamp,
};

axios.post(endpoint, payload)
  .then(response => {
    console.log('Timestamp sent successfully:', timestamp);
    console.log('UUID:', uuid);
    console.log('Server response:', response.data);
  })
  .catch(error => {
    console.error('Failed to send timestamp:', error.message);
  });