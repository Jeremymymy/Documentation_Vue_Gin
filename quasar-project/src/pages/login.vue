<template>
<div class="row justify-center full-height full-width text-center">
    <div class="column q-pa-lg">
      <div class="row">
        <q-card square dark class="q-pa-md q-ma-none no-shadow bg-grey-10" style="width:360px;">
          <q-card-section class="q-mt-xl q-mb-md">
          <q-btn round>
      <q-avatar rounded >
       <img src="~assets/tsmc.png">
      </q-avatar>
      </q-btn>
      <p class="text-h5 text-white">TSMC Document System</p>
            <p class="text-h5 text-white">Login </p>
          </q-card-section>
          <q-card-section>
            <q-form class="q-gutter-md">
              <q-input dark dense square filled clearable v-model="email" type="email" label="Email">
                <template v-slot:prepend>
                  <q-icon name="email" />
                </template>
              </q-input>
              <q-input dark dense square filled clearable v-model="password" type="password" label="Password">
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
              </q-input>
            </q-form>
          </q-card-section>
          <q-card-actions>
            <div class="row full-width items-center">
              <div class="col-6">
                <q-btn outline rounded size="mid" color="red-4" class="full-width text-white" label="Sign In" />
              </div>
               <div class="col-6">
                <q-btn outline rounded size="mid" color="red-4" class="full-width text-white" label="register" />
              </div>
            </div>
          </q-card-actions>
        </q-card>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      email: '',
      username: '',
      password: ''
    }
  }
}
</script>

<style>
</style>
