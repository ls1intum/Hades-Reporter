require('dotenv').config();
const axios = require('axios');
const dayjs = require('dayjs');

// 读取 .env 中的 ENDPOINT
const endpoint = process.env.ENDPOINT;

if (!endpoint) {
  console.error('❌ ENDPOINT is not defined in .env');
  process.exit(1);
}

// 格式化时间为 YYYY-MM-DD HH:mm:ss
const timestamp = dayjs().format('YYYY-MM-DD HH:mm:ss');

// 构造请求体
const payload = {
  timestamp: timestamp,
};

axios.post(endpoint, payload)
  .then(response => {
    console.log('✅ Timestamp sent successfully:', timestamp);
    console.log('🔁 Server response:', response.data);
  })
  .catch(error => {
    console.error('❌ Failed to send timestamp:', error.message);
  });
