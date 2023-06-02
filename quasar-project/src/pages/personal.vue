<template>
                  <q-form @submit.prevent="getAllUserInfo" class="q-gutter-md">
                    <q-btn label="確認" color="black" type="submit"/>
                  </q-form>
    <div  class="text-h6 q-pa-sm " align="center"><b class="title"><br/>個人資料</b></div>

    <div class="q-pa-md row section-card items-start q-gutter-md row justify-center"  align="center">
    <q-card class="my-card-info "  >
      <q-card-section horizontal>
        <q-img class="col-5" src="https://cdn.quasar.dev/img/boy-avatar.png"/>

        <q-card-section>
          <div class="col-5 ">
            <div class="row">
            <div class="col">人員名稱: </div>
            <div class="col" align="left">{{ userName }}</div>
            </div>
            <div class="row" >
            <div class="col" >所屬部門: </div>
            <div class="col" align="left">{{ userDep }}</div>
            </div>
            <div class="row">
            <div class="col">電子郵件: </div>
            <div class="col" align="left">{{ userEmail }}</div>
            </div>
        </div>
        </q-card-section>
      </q-card-section>

      <q-separator />
      <q-card-actions align="right">
              <!-- <div class="justify-center"> -->
              <q-btn flat color="black" label="修改聯絡資料" @click="dialog = true">
              <q-dialog v-model="dialog">
                <q-card class="my-card-info-fix q-pa-md">
                  <q-form @submit.prevent="modify" class="q-gutter-md">
                  <q-card-section>
                    <div class="text-h6">修改個人資料</div>
                  </q-card-section>
                  <q-card-section class="items-center q-gutter-sm">
                    <!-- <div class="row"> -->
                      <div class="col">姓名</div>
                      <!-- <div class="col"> -->
                      <q-input clearable filled color="balck" v-model="nameM" type="text" label="reset  name" />
                      <!-- </div> -->
                      <div class="col">電子郵件</div>
                      <!-- <div class="col"> -->
                      <q-input clearable filled color="balck" v-model="emailM" type="email" label="reset email"/>
                      <!-- </div> -->
                      <div class="col">密碼</div>
                      <!-- <div class="col"> -->
                      <q-input clearable filled color="balck" v-model="passwordM" type="text" label="reset password" />
                      <!-- </div> -->
                  </q-card-section>
                  <q-card-section class="row items-center q-gutter-sm justify-center">
                    <q-btn label="確認" color="black" type="submit"/>
                    <q-btn v-close-popup label="取消" color="black" />
                  </q-card-section>
                  </q-form>
                </q-card>
            </q-dialog>
            </q-btn>
      </q-card-actions>
    </q-card>
    </div>

    <div  class="text-h6 q-pa-sm " align="center"><b class="title"><br/>發布文章</b></div>
    <!-- <q-page-sticky position="top-right" :offset="[18, 18]"> -->
    <div class="column justify-evenly">
    <q-form
      @submit="createDoc"
      class="q-gutter-md"
    >
    <div class="col ">
    <q-input v-model="dName" label="文章標題" :dense="dense" class="section-card" />
    </div>
    <br>
    <div class="col ">
    <q-editor v-model="dContent" class="section-card"
        :toolbar="[
          ['bold', 'italic', 'strike', 'underline'],
          ['hr', 'link'],
          ['fullscreen'],
        ]"
    />
    </div>
    <!-- <div class="col "> -->
    <div class="row justify-center ">
    <q-btn @click="model = true"  type="submit">送出</q-btn>
    <div v-if="showSuccessMessage" class="success-message">
      <p class="message-text">Create successfully !</p>
    </div>
    </div>
    </q-form>
    </div>

    <!-- <div class = "row justify-center q-pa-md" align="center"> -->
    <div  class="text-h6 q-pa-sm " align="center"><b class="title">發文紀錄</b></div>
    <div class=" q-pa-md row section-card items-start q-gutter-md row justify-center"  align="center">
      <q-card class="my-card" flat bordered>
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        />
        <q-card-section>
          <div class="text-overline text-orange-9">Overline</div>
          <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
          <div class="text-caption text-grey">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn flat round color="red" icon="favorite" />
          <!-- <q-btn flat round color="teal" icon="bookmark" />
          <q-btn flat round color="primary" icon="share" /> -->
        </q-card-actions>
      </q-card>
       <q-card class="my-card" flat bordered>
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        />

        <q-card-section>
          <div class="text-overline text-orange-9">Overline</div>
          <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
          <div class="text-caption text-grey">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat round color="gray" icon="favorite" />
          <!-- <q-btn flat round color="teal" icon="bookmark" />
          <q-btn flat round color="primary" icon="share" /> -->
        </q-card-actions>
      </q-card>
       <q-card class="my-card" flat bordered>
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        />

        <q-card-section>
          <div class="text-overline text-orange-9">Overline</div>
          <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
          <div class="text-caption text-grey">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat round icon ="favorite" :color="expanded ? 'gray' : 'red'" @click="expanded = !expanded" />
          <!-- <q-btn flat round color="teal" icon="bookmark" />
          <q-btn flat round color="primary" icon="share" /> -->
        </q-card-actions>
      </q-card>
       <q-card class="my-card" flat bordered>
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        />

        <q-card-section>
          <div class="text-overline text-orange-9">Overline</div>
          <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
          <div class="text-caption text-grey">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat round color="gray" icon="favorite" />
          <!-- <q-btn flat round color="teal" icon="bookmark" />
          <q-btn flat round color="primary" icon="share" /> -->
        </q-card-actions>
      </q-card>
      <div class="q-pa-lg flex flex-center">
        <q-pagination
          v-model="current"
          :max="5"
          input
        />
      </div>
    </div>

   <!-- </div> -->
    <div  class="text-h6 q-pa-sm " align="center"><b class="title">收藏文件</b></div>
    <div class=" q-pa-md row section-card items-start q-gutter-md row justify-center"  align="center">
      <q-card class="my-card" flat bordered>
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        />
        <q-card-section>
          <div class="text-overline text-orange-9">Overline</div>
          <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
          <div class="text-caption text-grey">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </div>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn flat round color="red" icon="favorite" />
          <!-- <q-btn flat round color="teal" icon="bookmark" />
          <q-btn flat round color="primary" icon="share" /> -->
        </q-card-actions>
      </q-card>
       <q-card class="my-card" flat bordered>
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        />

        <q-card-section>
          <div class="text-overline text-orange-9">Overline</div>
          <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
          <div class="text-caption text-grey">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat round color="red" icon="favorite" />
          <!-- <q-btn flat round color="teal" icon="bookmark" />
          <q-btn flat round color="primary" icon="share" /> -->
        </q-card-actions>
      </q-card>
       <q-card class="my-card" flat bordered>
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        />

        <q-card-section>
          <div class="text-overline text-orange-9">Overline</div>
          <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
          <div class="text-caption text-grey">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat round color="red" icon="favorite" />
          <!-- <q-btn flat round color="teal" icon="bookmark" />
          <q-btn flat round color="primary" icon="share" /> -->
        </q-card-actions>
      </q-card>
       <q-card class="my-card" flat bordered>
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        />

        <q-card-section>
          <div class="text-overline text-orange-9">Overline</div>
          <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
          <div class="text-caption text-grey">
            {{ dName }}
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat round color="red" icon="favorite" />
          <!-- <q-btn flat round color="teal" icon="bookmark" />
          <q-btn flat round color="primary" icon="share" /> -->
        </q-card-actions>
      </q-card>
      <div class="q-pa-lg flex flex-center">
        <q-pagination
          v-model="current"
          :max="5"
          input
        />
      </div>
      </div>
      <!-- <vue-horizontal responsive>
      <section v-for="item in items" :key="item.title">
        <h3>{{ item.title }}</h3>
        <p>{{ item.content }}</p>
      </section>
      </vue-horizontal> -->
      <div  class="text-h6 q-pa-sm " align="center"><b class="title">TEST</b></div>
      <div class=" q-pa-md row section-card items-start q-gutter-md row justify-center"  align="center">
      <q-card class="my-card" flat bordered  v-for="item in allDoc" :key="item.ID">
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        />

        <q-card-section>
          <div class="text-overline text-orange-9">Overline</div>
          <div class="text-h5 q-mt-sm q-mb-xs">{{ item.Title}}</div>
          <div class="text-caption text-grey">
            {{ item.Content }}
          </div>
        </q-card-section>
        <q-card-actions align="right">
           <q-btn flat round icon="favorite" :color="item.favorite ? 'red' : 'gray'" @click="createCollect(item)"><!--:color="examColour" @click="item.favorite = !item.favorite" -->
          </q-btn>
        </q-card-actions>
      </q-card>

      <div class="q-pa-lg flex flex-center">
        <q-pagination
          v-model="current"
          :max="5"
          input
        />
      </div>
      </div>

</template>

<script>
import { ref } from 'vue'
import { useUserStore } from 'src/stores/user';
import axios from 'axios';
// import { useQuasar } from 'quasar'
import { LocalStorage, SessionStorage } from 'quasar';
// import { useQuasar } from 'quasar';
const userInfo = useUserStore();

export default {
  name: 'personal',
  setup () {
    const value = SessionStorage.getItem('userSession');
    const value2 = LocalStorage.getItem('userInfo');

    console.log(value);
    return {
      expanded: ref(false),
      slide: ref('style'),
      autoplay: ref(true),
      lorem: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
      dialog: ref(false),

      items: ref(),

      exambutton: ref(false),

      userName: ref(value2.Name),
      userDep: ref(value2.Department),
      userEmail: ref(value2.Email),
      userPassword: ref(value2.Password),
      userID: ref(value2.EmployeeId),

      nameM: ref(value2.Name),
      emailM: ref(value2.Email),
      passwordM: ref(value2.Password),

      dName: ref('title'),
      dContent: ref('content'),

      showSuccessMessage: false,

      allDoc: ref(''),
      allCollect: ref('')
    }
  },
  computed: {
    // examColour () {
    //   this.favorite = !this.favorite
    //   // return this.items.favorite === false ? 'red' : 'gray'
    // }
  },
  methods: {
    createCollect (ff) {
      ff.favorite = !ff.favorite;
      if (ff.favorite === true) {
        axios
          .get(`http://localhost:8000/TSMC/docs/collectDoc/${ff.ID}`)
          .then(response => {
            console.log(response);
            this.allCollect = response.data.Collect
            console.log(this.allCollect);
          })
          .catch(error => {
            console.error(error);
            // Handle registration error
            // You can display an error message or perform other actions
          });
      } else {
        console.log('Delete');
        axios
          .delete(`http://localhost:8000/TSMC/docs/deleteCollect/${ff.ID}`)
          .then(response => {
            console.log(response);
          })
          .catch(error => {
            console.error(error);
            // Handle registration error
            // You can display an error message or perform other actions
          });
      }
    },
    getAllUserInfo () {
      axios
        .get('http://localhost:8000/TSMC/users/getMyDetail')
        .then(response => {
          console.log(response);
          this.allDoc = response.data.PostDocs;
          this.allCollect = response.data.CollectDocs;
          // this.items = response.data.PostDocs;
          this.allDoc.forEach((elem) => {
            elem.favorite = false;
          });

          console.log(this.allDoc);
          console.log(this.allCollect);
          // userStore.modify({ user: response.data, name: response.data.Name, email: response.data.Email, password: response.data.Password });
          // Handle successful registration
          // LocalStorage.set('userInfo', response.data)
          // this.showSuccessMessage = true;
          // window.location.reload();
          // You can redirect the user or show a success message
        })
        .catch(error => {
          console.error(error);
          // Handle registration error
          // You can display an error message or perform other actions
        });
    },
    // deleteDoc () {
    //   axios
    //     .put(`http://localhost:8000/TSMC/users/${userInfo.getUserid}`, userM)
    //     .then(response => {
    //       console.log(response);
    //       console.log(value);

    //       userStore.modify({ user: response.data, name: response.data.Name, email: response.data.Email, password: response.data.Password });
    //       // Handle successful registration
    //       LocalStorage.set('userInfo', response.data)
    //       this.showSuccessMessage = true;
    //       window.location.reload();
    //       // You can redirect the user or show a success message
    //     })
    //     .catch(error => {
    //       console.error(error);
    //       // Handle registration error
    //       // You can display an error message or perform other actions
    //     });
    // },
    // getAllCollect () {
    //   axios
    //     .put(`http://localhost:8000/TSMC/users/${userInfo.getUserid}`, userM)
    //     .then(response => {
    //       console.log(response);
    //       console.log(value);

    //       userStore.modify({ user: response.data, name: response.data.Name, email: response.data.Email, password: response.data.Password });
    //       // Handle successful registration
    //       LocalStorage.set('userInfo', response.data)
    //       this.showSuccessMessage = true;
    //       window.location.reload();
    //       // You can redirect the user or show a success message
    //     })
    //     .catch(error => {
    //       console.error(error);
    //       // Handle registration error
    //       // You can display an error message or perform other actions
    //     });
    // },

    // deleteCollect () {
    //   axios
    //     .put(`http://localhost:8000/TSMC/users/${userInfo.getUserid}`, userM)
    //     .then(response => {
    //       console.log(response);
    //       console.log(value);

    //       userStore.modify({ user: response.data, name: response.data.Name, email: response.data.Email, password: response.data.Password });
    //       // Handle successful registration
    //       LocalStorage.set('userInfo', response.data)
    //       this.showSuccessMessage = true;
    //       window.location.reload();
    //       // You can redirect the user or show a success message
    //     })
    //     .catch(error => {
    //       console.error(error);
    //       // Handle registration error
    //       // You can display an error message or perform other actions
    //     });
    // },
    modify () {
      const userStore = useUserStore();
      const value = SessionStorage.getItem('userSession');

      console.log(value);
      // config.headers.Authorization = value;
      const userM = {
        id: this.userID,
        Name: this.nameM,
        Email: this.emailM,
        Password: this.passwordM
      };
      console.log(userM)
      axios
        .put(`http://localhost:8000/TSMC/users/${userInfo.getUserid}`, userM)
        .then(response => {
          console.log(response);
          console.log(value);

          userStore.modify({ user: response.data, name: response.data.Name, email: response.data.Email, password: response.data.Password });
          // Handle successful registration
          LocalStorage.set('userInfo', response.data)
          this.showSuccessMessage = true;
          window.location.reload();
          // You can redirect the user or show a success message
        })
        .catch(error => {
          console.error(error);
          // Handle registration error
          // You can display an error message or perform other actions
        });
    },
    createDoc () {
      const value = SessionStorage.getItem('userSession')
      console.log(value);
      // config.headers.Authorization = value;
      const doc = {
        Title: this.dName,
        Content: this.dContent
      };
      console.log(doc)
      axios
        .post('http://localhost:8000/TSMC/docs/createDoc', doc)
        .then(response => {
          console.log(response.data);
          console.log(value);
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

}
</script>

<style lang="sass" scoped>
  .section-card
    margin-right: 10%
    margin-left: 10%

  .my-card
    width: 100%
    height: 100%
    max-width: 300px
  .my-card-info
    height: 100%
    width: 100%
    max-width: 500px

  .my-card-info-fix
    width: 100%
    max-width: 400px
  .alignleft
    display: inline
    float: left
  .info-card
    margin-right: 10%
    margin-left: 10%
    margin-top:2%

  .center-in-center
    margin-top:2%
    left: 45%
  .success-message
    display: flex
    justify-content: center
    align-items: center
    background-color: green
    color: white
    padding: 10px
    margin-top: 10px

</style>
