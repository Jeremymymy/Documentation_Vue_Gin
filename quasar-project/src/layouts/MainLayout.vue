<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
    <q-toolbar class="bg-dark glossy text-white">
    <q-btn round>
      <q-avatar rounded >
       <img src="~assets/tsmc.png">
      </q-avatar>
      </q-btn>

      <q-toolbar-title>TSMC
      </q-toolbar-title>
      <q-btn flat round dense icon="search" />
      <div class="q-pa-md">
    <q-btn-dropdown
      class="glossy"
      color="dark"
      label="Account"
    >
    <div class="row no-wrap q-pa-md">
        <div class="column items-center">
          <q-btn color="dark" push
            size="sm"
            v-close-popup label ="登入" router-link to="/login"></q-btn><br/>
          <q-btn color="dark" push
            size="sm"
            v-close-popup label ="註冊" router-link to="/register"></q-btn>
          </div>
      </div>
    </q-btn-dropdown>
  </div>
    </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>

    <q-footer reveal elevated class="bg-dark glossy text-white">
      <q-toolbar>
          <q-avatar rounded>
            <img src="~assets/tsmc.png">
          </q-avatar>
          <q-toolbar-title>
          TSMC
        </q-toolbar-title>
      </q-toolbar>
    </q-footer>

  </q-layout>
</template>
