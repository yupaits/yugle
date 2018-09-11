import request from '../request'

const picturePathPrefix = '/picture';

const picture = {
  getBingPictures(page, size) {
    return request.api.get(`${picturePathPrefix}/bing?page=${page}&size=${size}`);
  },
  getShotOnOnePlusPictures(page, size) {
    return request.api.get(`${picturePathPrefix}/shot_on_oneplus?page=${page}&size=${size}`);
  }
};

export default picture