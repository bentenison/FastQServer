// import * as types from '../types'
import axios from "axios";
import router from "../../router";

export const state = {
  company: null,
  branch: null,
};
export const getters = {};
export const mutations = {};

export const actions = {
  //------------------>>>>>>>>>>>--[BRANCH ACTIONS]--<<<<<<<<<<<<<---------------------
  GET_BRANCH: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post("/branch/getbranch", payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  ADD_BRANCH: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post("/branch/addbranch", payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  UPDATE_BRANCH: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post("/branch/updatebranch", payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  DELETE_BRANCH: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .get(`/branch/deletebranch/${payload}`)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  GETALL_BRANCH: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .get(`/branch/getAllBranch/${payload}`)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  //------------------>>>>>>>>>>>--[USER ACTIONS]--<<<<<<<<<<<<<----------------------
  ADD_USER: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post(`/user/adduser`, payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  GETALL_USER: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post(`/user/getAlluser`, payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  GET_USER: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post(`/user/getuser`, payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  UPDATE_USER: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post(`/user/updateuser`, payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  DELETE_USER: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .get(`/user/deleteuser/${payload}`)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  //------------------>>>>>>>>>>>--[SERVICE ACTIONS]--<<<<<<<<<<<<<--------------------
  ADD_SERVICE: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post("/service/addservice", payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  GET_SERVICE: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post("/service/getservice", payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  UPDATE_SERVICE: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post("/service/updateservice", payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  DELETE_SERVICE: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .get(`/service/deleteservice/${payload}`)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  GETALL_SERVICES: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .get(`/service/getAllservice/${payload}`)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  //------------------>>>>>>>>>>>--[COUNTER ACTIONS]--<<<<<<<<<<<<<---------------------
  ADD_COUNTER: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post("/counter/addcounter", payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  GET_COUNTER: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post("/counter/getcounter", payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  UPDATE_COUNTER: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .post("/counter/updatecounter", payload)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  GETALL_COUNTERS: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .get(`"/counter/getAllcounter/${payload}`)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
  DELETE_COUNTER: ({ commit }, payload) => {
    return new Promise((resolve, reject) => {
      axios
        .get(`/counter/deletecounter/${payload}`)
        .then((res) => {
          console.log("res:::::", res);
          resolve(res.data);
          // this.$toast.success("service added successfully.");
        })
        .catch((err) => {
          console.log(err.response);
          reject(err);
          // this.$toast.error("error occured while adding branch!!!");
        });
    });
  },
};
