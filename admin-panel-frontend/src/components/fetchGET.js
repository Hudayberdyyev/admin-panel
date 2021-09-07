const url = "http://backend:7770";

export const get = (directory) => fetch(`${url}${directory}`,{
    method: 'GET',
    headers:{
        Authorization: "Bearer " + sessionStorage.getItem('token')
    }
});