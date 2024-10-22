import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 10,  // количество виртуальных пользователей (VUs)
    duration: '30s',  // время нагрузки
};


export default function(){
    const res = http.get('http://localhost:8080'); // Url

    check(res, {
        'status is 200': (r) => r.status === 200,
    });

    sleep(1);
}