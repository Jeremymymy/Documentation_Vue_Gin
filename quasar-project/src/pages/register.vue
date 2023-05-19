<template>
  <div class="row justify-center full-height full-width text-center">
    <div class="column q-pa-lg">
      <div class="row">
        <q-card square dark class="q-pa-md q-ma-none no-shadow bg-grey-10" style="width:360px;">
          <q-card-section class="q-mt-xl q-mb-md">
            <q-btn round>
              <q-avatar rounded>
                <img src="~assets/tsmc.png">
              </q-avatar>
            </q-btn>
            <p class="text-h5 text-white">TSMC Document System</p>
            <p class="text-h5 text-white">Register</p>
          </q-card-section>
          <q-card-section>
            <q-form class="q-gutter-md" @submit.prevent="register">
              <q-input dark dense square filled clearable v-model="employeeId" type="employeeId" label="Employee_Id">
                <template v-slot:prepend>
                  <q-icon name="person" />
                </template>
              </q-input>
              <q-input dark dense square filled clearable v-model="email" type="email" label="Email">
                <template v-slot:prepend>
                  <q-icon name="email" />
                </template>
              </q-input>
              <q-input dark dense square filled clearable v-model="username" type="text" label="Username">
                <template v-slot:prepend>
                  <q-icon name="person" />
                </template>
              </q-input>
              <q-input dark dense square filled clearable v-model="password" type="password" label="Password">
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>
              <q-select dark dense square filled clearable v-model="section" :options="options" label="Department">
                <template v-slot:prepend>
                  <q-icon name="work" />
                </template>
              </q-select>
              <div class="row full-width items-center">
                <div class="col-6">
                  <q-btn outline rounded size="mid" color="red-4" class="full-width text-white" label="Register" type="submit" />
                </div>
                <div class="col-6">
                  <q-btn outline rounded size="mid" color="red-4" class="full-width text-white" label="Sign in" router-link to="/login" />
                </div>
                <div v-if="showSuccessMessage" class="success-message">
                  <p class="message-text">Register successful ! Now you can sign in</p>
                </div>
              </div>
            </q-form>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Login',
  data () {
    return {
      showSuccessMessage: false,
      employeeId: '',
      email: '',
      username: '',
      password: '',
      section: null,
      options: ['HR', 'Sales', 'Marketing', 'IT', 'Finance']
    };
  },
  methods: {
    register () {
      const user = {
        EmployeeId: this.employeeId,
        Name: this.username,
        Email: this.email,
        Password: this.password,
        Department: this.section
      };
      console.log(user)
      axios
        .post('http://localhost:8000/TSMC/users/', user)
        .then(response => {
          console.log(response.data);
          // Handle successful registration
          this.showSuccessMessage = true;
          // You can redirect the user or show a success message
        })
        .catch(error => {
          console.error(error);
          // Handle registration error
          // You can display an error message or perform other actions
        });
    }
  }
};
</script>

<style scoped>
.success-message {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: green;
  color: white;
  padding: 10px;
  margin-top: 10px;
}

.message-text {
  margin: 0;
  font-size: 16px;
}
</style>
