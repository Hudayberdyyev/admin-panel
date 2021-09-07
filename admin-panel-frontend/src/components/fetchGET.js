const url = "http://144.91.104.118:7770";

export const get = (directory) => fetch(`${url}${directory}`,{
    method: 'GET',
    headers:{
        Authorization: "Bearer " + sessionStorage.getItem('token')
    }
});