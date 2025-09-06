import axios from 'axios';

export const fetchAll = (baseURL, tokenKey) => {
  return new Promise((resolve, reject) => {
    const token = localStorage.getItem(tokenKey);
    axios.get(baseURL + 'api/v1/document-types', {
      params: {},
      headers:{
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      }
    }).then(function (response) {
      resolve(response);
    }).catch(function (error) {
      if(error.response.status == 404){
        console.error('Miembro no encontrado')
      }else{
        console.error(error.response.data);
      }
      reject(error.response);
    })
    .then(function () {
      // todo?
    });
  });
}