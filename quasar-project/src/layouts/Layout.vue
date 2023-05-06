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
      <q-btn flat class="q-mr-md" label ="發表文件" ></q-btn>
      <q-btn flat class="q-mr-md" label ="管理文件" ></q-btn>
      <q-btn flat class="q-mr-md" label ="收藏文件" ></q-btn>
      <q-btn flat round dense icon="search" />
      <div class="q-pa-md">
    <q-btn-dropdown
      class="glossy"
      color="dark"
      label="Account"
    >
    <div class="row no-wrap q-pa-md">
        <div class="column items-center">
          <q-avatar size="72px">
            <img src="https://cdn.quasar.dev/img/boy-avatar.png">
          </q-avatar>

          <div class="text-subtitle1 q-mt-md q-mb-xs">John Doe</div>
          <q-btn
            color="dark"
            label="基本資料"
            push
            size="sm"
            v-close-popup
          /><br/>
          <q-btn
            color="dark"
            label="Logout"
            push
            size="sm"
            v-close-popup
          />
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
