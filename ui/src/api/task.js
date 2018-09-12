import request from '../request'

const taskPathPrefix = '/task';

const task = {
  getTasks(page, size) {
    return request.api.get(`${taskPathPrefix}/page?page=${page}&size=${size}`);
  }
};

export default task