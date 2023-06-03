<template>
    <div  class="text-h6 q-pa-sm " align="center"><b class="title"><br/>個人資料</b></div>

    <div class="q-pa-md row section-card items-start q-gutter-md row justify-center"  align="center">
    <q-card class="my-card-info "  >
      <q-card-section horizontal>
        <q-img class="col-5" src="https://cdn.quasar.dev/img/boy-avatar.png"/>
        <!-- https://cdn.quasar.dev/img/avatar.png -->
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

      <div  class="text-h6 q-pa-sm " align="center"><b class="title">發文紀錄</b></div>
      <div class=" q-pa-md row section-card items-start q-gutter-md justify-center"  align="center">
      <q-card class="my-card col" flat bordered  v-for="item in paginatedDoc()" :key="item.ID">
        <!-- <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        /> -->
        <q-card-actions align="left">
           <q-btn flat round icon="favorite" :color="item.favorite ? 'red' : 'gray'" @click="createCollect(item)"/>
        </q-card-actions>

        <q-card-section>

          <div class="text-h5 q-mt-sm q-mb-xs">{{ item.Title }}</div>
          <div class="text-overline text-orange-10">Author: {{ item.AuthorName }}</div>
          <div class="text-caption text-grey">
            {{ item.Content }}
          </div>
        </q-card-section>
        <q-card-actions align="right">
           <!-- <q-btn flat round icon="favorite" :color="item.favorite ? 'red' : 'gray'" @click="createCollect(item)"/> -->
            <router-link :to="{path: '/detail', query: {docID: item.ID }}">
              <q-btn flat round color="primary" icon="edit" />
            </router-link>

            <q-btn flat round color="teal" icon="delete" @click="item.deleteDialog = !item.deleteDialog" >
              <q-dialog v-model="item.deleteDialog">
                <q-card class="my-card-info-fix q-pa-md" align="center">
                  <q-card-section>
                    <div class="text-h6">確認要刪除 {{ item.Title }} ?</div>
                  </q-card-section>

                <q-btn flat round color="black" label="確認" type="submit" @click="deleteDoc(item)" ></q-btn>
                <q-btn flat round v-close-popup label="取消" color="black"></q-btn>
                </q-card>
              </q-dialog>
            </q-btn>
        </q-card-actions>
      </q-card>
      <q-dialog v-model="dialogDelete">
                <!-- <q-card class="my-card-info-fix q-pa-md" align="center"> -->
                  <q-card-section>
                    <div class="text-h6">確認要刪除 {{ item.Title }} ?</div>
                  </q-card-section>

                <q-btn flat round color="black" label="確認" type="submit" @click="deleteDoc(item)" ></q-btn>
                <q-btn flat round v-close-popup label="取消" color="black"></q-btn>
                <!-- </q-card> -->
      </q-dialog>
      </div>
      <div class="q-pa-lg flex flex-center justify-center">
        <q-pagination
          v-model="currentDoc"
          :min="1"
          :max="Math.ceil(allDoc.length/4)"
          :input="true"
          input-class="text-orange-10"
        />
      </div>

      <div  class="text-h6 q-pa-sm " align="center"><b class="title">收藏文件</b></div>
      <div class=" q-pa-md section-card items-start q-gutter-md row justify-center"  align="center">
      <!-- <div class="q-pa-lg center row"> -->
        <q-card class="my-card col" flat bordered  v-for="item in paginatedCol()" :key="item.ID">
          <!-- <q-img
            src="https://cdn.quasar.dev/img/parallax2.jpg"
          /> -->
          <q-card-actions align="left">
            <q-btn flat round icon="favorite" color= 'red' @click="deleteCollect(item)"/>
          </q-card-actions>

          <q-card-section>
            <div class="text-h5 q-mt-sm q-mb-xs">{{ item.Title}}</div>
            <div class="text-overline text-orange-10">Author: {{ item.AuthorName }}</div>
            <div class="text-caption text-grey">
              {{ item.Content }}
            </div>
          </q-card-section>
          <q-card-actions align="right">
            <!-- <q-btn flat round color="primary" icon="edit" /> -->
            <!-- <q-btn flat round color="teal" icon="delete" /> -->
            <router-link :to="{path: '/detail', query: {docID: item.DocId }}">
              <q-btn flat round color="primary" icon="edit" />
            </router-link>
          </q-card-actions>
        </q-card>
      </div>
      <!-- </div> -->
      <!-- <q-card-section> -->
      <div class="q-pa-lg center row justify-center">
        <q-pagination
          v-model="currentCollect"
          :min="1"
          :max="Math.ceil(allCollect.length/4)"
          :input="true"
          input-class="text-orange-10"
        />
      </div>
      <!-- </q-card-section> -->
      <!-- </div> -->

</template>

<script>
import { ref } from 'vue'
import { useUserStore } from 'src/stores/user';
import axios from 'axios';
// import { useQuasar } from 'quasar'
import { LocalStorage, SessionStorage } from 'quasar';
// import { useQuasar } from 'quasar';
// const userInfo = useUserStore();
export default {
  name: 'personal',
  setup () {
    const value = SessionStorage.getItem('userSession');
    const value2 = LocalStorage.getItem('userInfo');
    const expanded = ref(false);
    const dialog = ref(false);

    const userName = ref(value2.Name);
    const userDep = ref(value2.Department);
    const userEmail = ref(value2.Email);
    const userPassword = ref(value2.Password);
    const userID = ref(value2.EmployeeId);

    const nameM = ref(value2.Name);
    const emailM = ref(value2.Email);
    const passwordM = ref(value2.Password);

    const dName = ref('title');
    const dContent = ref('content');

    let showSuccessMessage = false;

    const allDoc = ref('');
    const allCollect = ref('');
    const currentCollect = ref(1);
    const currentDoc = ref(1);
    const pageSize = 4;

    function totalPages (item) {
      return Math.ceil(item.length / pageSize);
    };
    function paginatedDoc () {
      const startIndex = (currentDoc.value - 1) * pageSize;
      return allDoc.value.slice(startIndex, startIndex + pageSize);
    };
    function paginatedCol () {
      const startIndex = (currentCollect.value - 1) * pageSize;
      return allCollect.value.slice(startIndex, startIndex + pageSize);
    };
    function createCollect (ff) {
      ff.favorite = !ff.favorite;
      if (ff.favorite === true) {
        axios
          .get(`http://localhost:8000/TSMC/docs/collectDoc/${ff.ID}`)
          .then(response => {
            console.log(response);
            allCollect.value = response.data.Collect
            console.log(allCollect.value);
            getAllUserInfo();
          })
          .catch(error => {
            console.error(error);
            // Handle registration error
          });
      } else {
        console.log('Delete');
        let colID = null;
        allCollect.value.forEach((elem) => {
          if (ff.ID === elem.DocId) {
            colID = elem.ID
          }
        });
        if (colID !== null) {
          axios
            .delete(`http://localhost:8000/TSMC/docs/deleteCollect/${colID}`)
            .then(response => {
              console.log(response);
              getAllUserInfo();
            })
            .catch(error => {
              console.error(error);
            });
        }
      }
    };
    function getAllUserInfo () {
      axios
        .get('http://localhost:8000/TSMC/users/getMyDetail')
        .then(response => {
          console.log(response);

          allDoc.value = response.data.PostDocs;
          allCollect.value = response.data.CollectDocs;
          console.log(allCollect.value.length);
          allDoc.value.forEach((elem) => {
            // elem.page =
            allCollect.value.forEach((elemC) => {
              // console.log(elem);
              // console.log(elemC.Title);
              if (elemC.DocId !== elem.ID && elem.favorite !== true) { /* && elemC.Title === elem.Title */
                elem.favorite = false;
                // console.log('yes');
              } else {
                elem.favorite = true;
                // console.log('no');
              }
              elem.deleteDialog = false;
            });
          });
          LocalStorage.set('userCollect', allCollect.value);
          console.log(allDoc.value);
          console.log(allCollect.value);
        })
        .catch(error => {
          console.error(error);
          // Handle registration error
          // You can display an error message or perform other actions
        });
    };
    function deleteCollect (ff) {
      // ff.favorite = !ff.favorite;
      console.log('Delete');
      axios
        .delete(`http://localhost:8000/TSMC/docs/deleteCollect/${ff.ID}`)
        .then(response => {
          console.log(response);
          getAllUserInfo();
        })
        .catch(error => {
          console.error(error);
          // Handle registration error
          // You can display an error message or perform other actions
        });
    };
    function modify () {
      const userStore = useUserStore();
      const value = SessionStorage.getItem('userSession');
      const value2 = LocalStorage.getItem('userInfo');

      console.log(value);
      // config.headers.Authorization = value;
      const userM = {
        id: userID.value,
        Name: nameM.value,
        Email: emailM.value,
        Password: passwordM.value
      };
      console.log(userM)

      if (userM.Name === null) {
        userM.Name = value2.Name
      } else if (userM.Email === null) {
        userM.Email = value2.Email
      } else if (userM.Password === null) {
        userM.Password = value2.Password
      }

      axios
        .put(`http://localhost:8000/TSMC/users/${value2.EmployeeId}`, userM)
        .then(response => {
          console.log(response);
          console.log(value);
          userStore.modify({ user: response.data, name: response.data.Name, email: response.data.Email, password: response.data.Password });
          response.data.Password = passwordM.value;
          LocalStorage.set('userInfo', response.data)
          showSuccessMessage = true;
          window.location.reload();
        })
        .catch(error => {
          console.error(error);
        });
    };
    function createDoc () {
      const value = SessionStorage.getItem('userSession')
      console.log(value);
      // config.headers.Authorization = value;
      const doc = {
        Title: dName.value,
        Content: dContent.value
      };
      let tmpTitle = doc.Title;
      allDoc.value.forEach((elem) => {
        if (elem.Title === doc.Title || elem.Title.includes(doc.Title + ' (')) {
          if (elem.Title.includes(' (') && elem.Title.includes(')')) {
            // console.log('abc ' + elem)
            const split = elem.Title.split('(')
            // console.log('length ' + split.length)
            // console.log('length ' + split[split.length - 1])
            let index = parseInt(split[split.length - 1], 10);
            const substr = `(${index})`;
            const substrNew = `(${++index})`;
            // console.log('substr ' + substr)
            // console.log('substrM ' + substrNew)
            tmpTitle = elem.Title.replace(substr, substrNew)
          } else {
            console.log('def')
            tmpTitle = doc.Title + ' (1)';
          }
        }
      });
      doc.Title = tmpTitle
      console.log(doc)
      axios
        .post('http://localhost:8000/TSMC/docs/createDoc', doc)
        .then(response => {
          console.log(response.data);
          console.log(value);
          // Handle successful registration
          showSuccessMessage = true;
          getAllUserInfo();
        })
        .catch(error => {
          console.error(error);
          // Handle registration error
          // You can display an error message or perform other actions
        });
    };
    function deleteDoc (ff) {
      axios
        .delete(`http://localhost:8000/TSMC/docs/deleteDoc/${ff.ID}`)
        .then(response => {
          console.log(response);
          console.log(response.data);
          allCollect.value.forEach((elemC) => {
            if (ff.AuthorId === elemC.AuthorId && ff.Title === elemC.Title) {
              deleteCollect(elemC);
            }
          });
          getAllUserInfo();
        })
        .catch(error => {
          console.error(error);
        });
    };
    function updateDoc (ff) {
      // router.push('/detail');
      // axios
      //   .delete(`http://localhost:8000/TSMC/docs/deleteDoc/${ff.ID}`)
      //   .then(response => {
      //     console.log(response);
      //     console.log(response.data);
      //     allCollect.value.forEach((elemC) => {
      //       if (ff.AuthorId === elemC.AuthorId && ff.Title === elemC.Title) {
      //         deleteCollect(elemC);
      //       }
      //     });
      //     getAllUserInfo();
      //   })
      //   .catch(error => {
      //     console.error(error);
      //   });
    };

    getAllUserInfo();
    console.log(value);
    return {
      currentDoc,
      currentCollect,
      expanded,
      dialog,

      userName,
      userDep,
      userEmail,
      userPassword,
      userID,

      nameM,
      emailM,
      passwordM,

      dName,
      dContent,

      showSuccessMessage,

      allDoc,
      allCollect,
      getAllUserInfo,
      createCollect,
      deleteCollect,
      createDoc,
      modify,
      totalPages,
      paginatedDoc,
      paginatedCol,
      deleteDoc,
      updateDoc
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
