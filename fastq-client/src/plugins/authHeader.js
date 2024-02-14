import {getTokenPayload,getTokens,storeTokens,removeTokens} from "./utils"

export default function authHeader() {
    let user = JSON.parse(localStorage.getItem('user'));
  
    if (user && user.accessToken) {
      return { Authorization: 'Bearer ' + user.accessToken }; // for Spring Boot back-end
      // return { 'x-access-token': user.accessToken };       // for Node.js Express back-end
    } else {
      return {};
    }
  }

 export const authenticate = async (email, password, url) => {
    // state.isLoading = true;
    // state.error = null;
  
    // const { data, error } = await doRequest({
    //   url,
    //   method: 'post',
    //   data: {
    //     email,
    //     password,
    //   },
    // });
  
    // if (error) {
    //   state.error = error;
    //   state.isLoading = false;
    //   return;
    // }
  
    // const { tokens } = data;
  
    // storeTokens(tokens.idToken, tokens.refreshToken);
  
    // const tokenClaims = getTokenPayload(tokens.idToken);
  
    // // set tokens to local storage with expiry (separate function)
    // state.idToken = tokens.idToken;
    // state.currentUser = tokenClaims.user;
    // state.isLoading = false;
  };