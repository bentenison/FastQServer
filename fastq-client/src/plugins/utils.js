import jwt_decode from 'jwt-decode';

export const storeTokens = (idToken, refreshToken) => {
    sessionStorage.setItem("user-token", idToken);
    sessionStorage.setItem("refresh-token", refreshToken);
  };
  
  export const removeTokens = () => {
    sessionStorage.removeItem("user-token");
    sessionStorage.removeItem("refresh-token");
  };
  
  export const getTokens = () => {
    return [
      sessionStorage.getItem("user-token"),
      sessionStorage.getItem("refresh-token"),
    ];
  };
  
  // gets the token's payload, and returns null
  // if invalid
  export const getTokenPayload = (token) => {
    if (!token) {
      return null;
    }
  
    const tokenClaims = jwt_decode(token);
  
    if (Date.now() / 1000 >= tokenClaims.exp) {
      return null;
    }
    // console.log("claims::::::::",tokenClaims);
    return tokenClaims;
  };