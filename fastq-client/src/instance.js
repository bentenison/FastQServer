// axiosInstances.js

import axios from 'axios';

const localinstance = axios.create({
    baseURL: '/local',
});

// const instance2 = axios.create({
//   baseURL: 'https://api.example.com/v2',
// });

export { localinstance };
