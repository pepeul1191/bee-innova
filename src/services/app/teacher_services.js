import axios from 'axios';

export const create = (baseURL, tokenKey, data) => {
  delete data._id;
  return new Promise((resolve, reject) => {
    const token = localStorage.getItem(tokenKey);
    axios.post(baseURL + 'api/v1/teachers', data, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      }
    })
    .then(function (response) {
      resolve(response);
    })
    .catch(function (error) {
      if (error.response?.status === 404) {
        console.error('Miembro no encontrado');
      } else {
        console.error(error.response?.data || error.message);
      }
      reject(error.response);
    });
  });
}

export const update = (baseURL, tokenKey, data) => {
  const _id = data._id;
  delete data._id;
  return new Promise((resolve, reject) => {
    const token = localStorage.getItem(tokenKey);
    axios.put(baseURL + 'api/v1/teachers/' + _id, data, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      }
    })
    .then(function (response) {
      resolve(response);
    })
    .catch(function (error) {
      if (error.response?.status === 404) {
        console.error('Miembro no encontrado');
      } else {
        console.error(error.response?.data || error.message);
      }
      reject(error.response);
    });
  });
}

export const fetchOneById = (baseURL, tokenKey, issueId) => {
  return new Promise((resolve, reject) => {
    const token = localStorage.getItem(tokenKey);
    axios.get(baseURL + 'api/v1/teachers/fetch-one/' + issueId, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      }
    })
    .then(function (response) {
      resolve(response);
    })
    .catch(function (error) {
      if (error.response?.status === 404) {
        console.error('Miembro no encontrado');
      } else {
        console.error(error.response?.data || error.message);
      }
      reject(error.response);
    });
  });
}

export const editIssue = (baseURL, tokenKey, data, issueId) => {
  return new Promise((resolve, reject) => {
    const token = localStorage.getItem(tokenKey);
    axios.put(baseURL + 'api/v1/teachers/' + issueId, data, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      }
    })
    .then(function (response) {
      resolve(response);
    })
    .catch(function (error) {
      if (error.response?.status === 404) {
        console.error('Miembro no encontrado');
      } else {
        console.error(error.response?.data || error.message);
      }
      reject(error.response);
    });
  });
}

export const patchTags = (baseURL, tokenKey, issueId, data) => {
  return new Promise((resolve, reject) => {
    const token = localStorage.getItem(tokenKey);
    axios.patch(baseURL + 'api/v1/teachers/' + issueId , data, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      }
    })
    .then(function (response) {
      resolve(response);
    })
    .catch(function (error) {
      if (error.response?.status === 404) {
        console.error('Miembro no encontrado');
      } else {
        console.error(error.response?.data || error.message);
      }
      reject(error.response);
    });
  });
}

export const addDocument = (baseURL, tokenKey, issueId, data) => {
  return new Promise((resolve, reject) => {
    const token = localStorage.getItem(tokenKey);
    axios.post(baseURL + 'api/v1/teachers/' + issueId + '/documents', data, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      }
    })
    .then(function (response) {
      resolve(response);
    })
    .catch(function (error) {
      if (error.response?.status === 404) {
        console.error('Miembro no encontrado');
      } else {
        console.error(error.response?.data || error.message);
      }
      reject(error.response);
    });
  });
}
// api/v1/teachers/68619302d14017435b6d3316/documents/686196bffde75fcaf64413c8
export const removeDocument = (baseURL, tokenKey, issueId, documentId) => {
  return new Promise((resolve, reject) => {
    const token = localStorage.getItem(tokenKey);
    axios.delete(baseURL + 'api/v1/teachers/' + issueId + '/documents/' + documentId, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      }
    })
    .then(function (response) {
      resolve(response);
    })
    .catch(function (error) {
      if (error.response?.status === 404) {
        console.error('Miembro no encontrado');
      } else {
        console.error(error.response?.data || error.message);
      }
      reject(error.response);
    });
  });
}
