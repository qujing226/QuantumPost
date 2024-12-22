import axios from 'axios';

const API_URL = 'http://localhost:8080';

// 查看课程
export const getCourses = (number) => {
    return axios.get(`/api/courses`,{params:{
        number,
        }});
};

// 调课申请
export const requestChangeCourse = (courseChangeRequest) => {
    return axios.post(`${API_URL}/courses/change-request`, courseChangeRequest);
};
