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

    <div id="target1" class="text-h6 q-pa-sm " align="center"><b class="title"><br/>發布文章</b></div>
    <!-- <q-page-sticky position="top-right" :offset="[18, 18]"> -->
    <div class="column justify-evenly">
    <q-form
      @submit="createDoc"
      class="q-gutter-md"
    >
    <div class="col ">
    <q-input v-model="dName" label="文章標題" class="section-card" />
    </div>
    <div class="col">
      <q-select v-model="section" :options="options" label="Department" class="section-card">
        <template v-slot:prepend>
          <q-icon name="work" />
        </template>
      </q-select>
    </div>
    <br>
    <div class="col ">
    <q-editor v-model="dContent" class="section-card"
        :toolbar="[
          ['bold', 'italic', 'strike', 'underline'],
          ['fullscreen'],
        ]"
    />
    </div>
    <!-- <div class="col "> -->
    <div class="column  items-center">
    <div class="col">
    <q-btn @click="model = true"  type="submit">送出</q-btn>
    </div>
    <div v-if="showSuccessMessage" class="success-message col">
      <p class="message-text">Create successfully !</p>
    </div>
    </div>
    </q-form>
    </div>

      <div id="target2" class="text-h6 q-pa-sm " align="center"><b class="title">發文紀錄</b></div>
      <div class=" q-pa-md row section-card items-start q-gutter-md justify-center"  align="center">
      <q-card class="my-card col" flat bordered  v-for="item in paginatedDoc()" :key="item.ID">
        <!-- <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg"
        /> -->
        <q-card-actions align="between">
           <q-btn flat round icon="favorite" :color="item.favorite ? 'red' : 'gray'" @click="createCollect(item)"/>
           <q-btn flat round color="teal" icon="delete" @click="item.deleteDialog = !item.deleteDialog"/>
        </q-card-actions>

        <q-card-section>

          <div class="text-h5 q-mt-sm q-mb-xs">{{ item.Title }}</div>
          <div class="text-overline text-orange-10">Author: {{ item.AuthorName }}</div>
          <div class="text-overline text-orange-10">Department: {{ item.Belong }}</div>
          <div class="text-caption text-grey">
            {{ item.Content }}
          </div>
        </q-card-section>
        <q-card-actions align="center">
          <router-link :to="{path: '/detail', query: {docID: item.ID }}">
              <q-btn unelevated rounded glossy color="primary" icon="description" label = "文章內容"/>
            </router-link>

              <q-dialog v-model="item.deleteDialog">
                <q-card class="my-card-info-fix q-pa-md" align="center">
                  <q-card-section>
                    <div class="text-h6">確認要刪除 {{ item.Title }} ?</div>
                  </q-card-section>

                <q-btn flat round color="black" label="確認" type="submit" @click="deleteDoc(item)" ></q-btn>
                <q-btn flat round v-close-popup label="取消" color="black"></q-btn>
                </q-card>
              </q-dialog>
        </q-card-actions>
      </q-card>

      </div>
      <div class="q-pa-lg flex flex-center justify-center">
        <q-pagination
          v-model="currentDoc"
          :min="1"
          :max="totalPages(allDoc)"
          :input="true"
          input-class="text-orange-10"
        />
      </div>

      <div id="target3" class="text-h6 q-pa-sm " align="center"><b class="title">收藏文件</b></div>
      <div class=" q-pa-md section-card items-start q-gutter-md row justify-center"  align="center">
      <!-- <div class="q-pa-lg center row"> -->
        <q-card class="my-card col" flat bordered  v-for="item in paginatedCol()" :key="item.ID">
          <!-- <q-img
            src="https://cdn.quasar.dev/img/parallax2.jpg"
          /> -->
          <q-card-actions align="between">
            <q-btn flat round icon="favorite" color= 'red' @click="deleteCollect(item)"/>
            <!--<q-btn v-show="item.mine" flat round color="teal" icon="delete" @click="item.deleteDialog = !item.deleteDialog"/> -->
          </q-card-actions>

          <q-card-section>
            <div class="text-h5 q-mt-sm q-mb-xs">{{ item.Title}}</div>
            <div class="text-overline text-orange-10">Author: {{ item.AuthorName }}</div>
            <div class="text-overline text-orange-10">Department: {{ item.Belong }}</div>
            <div class="text-caption text-grey">
              {{ item.Content }}
            </div>
          </q-card-section>
          <q-card-actions align="center">
            <router-link :to="{path: '/detail', query: {docID: item.ID }}">
              <q-btn unelevated rounded glossy color="primary" icon="description" label = "文章內容"/>
            </router-link>
              <q-dialog v-model="item.deleteDialog">
                <q-card class="my-card-info-fix q-pa-md" align="center">
                  <q-card-section>
                    <div class="text-h6">確認要刪除 {{ item.Title }} ?</div>
                  </q-card-section>

                <q-btn flat round color="black" label="確認" type="submit" @click="deleteDoc(item)" ></q-btn>
                <q-btn flat round v-close-popup label="取消" color="black"></q-btn>
                </q-card>
              </q-dialog>
          </q-card-actions>
        </q-card>
      </div>
      <div class="q-pa-lg center row justify-center">
        <q-pagination
          v-model="currentCollect"
          :min="1"
          :max="totalPages(allCollect)"
          :input="true"
          input-class="text-orange-10"
        />
      </div>

</template>

<script>
import { ref, onBeforeUpdate } from 'vue'
import { useRoute } from 'vue-router';
// import { useUserStore } from 'src/stores/user';
import axios from 'axios';
import { LocalStorage, SessionStorage } from 'quasar';

export default {
  name: 'personal',
  setup () {
    const value = SessionStorage.getItem('userSession');
    const userInfo = LocalStorage.getItem('userInfo');
    const expanded = ref(false);
    const dialog = ref(false);

    const userName = ref(userInfo.Name);
    const userDep = ref(userInfo.Department);
    const userEmail = ref(userInfo.Email);
    const userPassword = ref(userInfo.Password);
    const userID = ref(userInfo.EmployeeId);

    const nameM = ref(userInfo.Name);
    const emailM = ref(userInfo.Email);
    const passwordM = ref(userInfo.Password);

    const dName = ref('title');
    const dContent = ref('content');

    const showSuccessMessage = ref(false);

    const allDoc = ref(Object([{}]));
    const allCollect = ref(Object([{}]));
    const currentCollect = ref(1);
    const currentDoc = ref(1);
    const pageSize = 4;

    const section = ref(userInfo.Department);
    const options = [userInfo.Department, 'Public'];

    function totalPages (item) {
      console.log(item)
      if (item.length === 0) {
        return 1;
      } else {
        return Math.ceil(item.length / pageSize);
      }
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
        let collectedData;
        axios
          .get(`http://localhost:8000/TSMC/docs/collectDoc/${ff.ID}`)
          .then(response => {
            console.log(response.data);
            collectedData = response.data.Collect;
            console.log(response.data.Collect);
            console.log(collectedData);
            collectedData.ID = response.data.Collect.DocId
            console.log(collectedData.ID);
            allCollect.value = collectedData
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
          const modifiedCollectDocs = response.data.CollectDocs.map(doc => {
            console.log(doc.ID);
            console.log(doc.DocId);
            doc.ID = doc.DocId; // 将元素中的 ID 值替换为 DocId 值
            return doc; // 返回修改后的元素
          });
          allDoc.value = response.data.PostDocs;
          allCollect.value = modifiedCollectDocs;
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

              if (elemC.AuthorId === userInfo.EmployeeId) {
                elemC.mine = true;
              } else {
                elemC.mine = false;
              }
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
        });
    };
    function modify () {
      // const userStore = useUserStore();
      // const value = SessionStorage.getItem('userSession');
      // const userInfo = LocalStorage.getItem('userInfo');
      console.log(value);

      const userM = {
        id: userID.value,
        Name: nameM.value,
        Email: emailM.value,
        Password: passwordM.value
      };
      console.log(userM)

      if (userM.Name === null) {
        userM.Name = userInfo.Name
      }
      if (userM.Email === null) {
        userM.Email = userInfo.Email
      }
      if (userM.Password === null) {
        userM.Password = userInfo.Password
      }

      axios
        .put(`http://localhost:8000/TSMC/users/${userInfo.EmployeeId}`, userM)
        .then(response => {
          console.log(response);
          console.log(value);
          // userStore.modify({ user: response.data, name: response.data.Name, email: response.data.Email, password: response.data.Password });
          response.data.Password = passwordM.value;
          LocalStorage.set('userInfo', response.data)
          // showSuccessMessage.value = true;
          window.location.reload();
        })
        .catch(error => {
          console.error(error);
        });
    };
    function createDoc () {
      const doc = {
        Title: dName.value,
        Content: dContent.value,
        Belong: section.value
      };
      let tmpTitle = doc.Title;
      allDoc.value.forEach((elem) => {
        if (elem.Title === doc.Title || elem.Title.includes(doc.Title + ' (')) {
          if (elem.Title.includes(' (') && elem.Title.includes(')')) {
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

          getAllUserInfo();
          showSuccessMessage.value = true;
          setTimeout(function close () { showSuccessMessage.value = !showSuccessMessage.value }, 2000);
        })
        .catch(error => {
          console.error(error);
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
    // function redirectToIndex () {
    //   // 登录成功，进行页面跳转
    //   this.$router.push('/detail');
    // }

    const $route = useRoute();
    if ($route.params.section != null) {
      onBeforeUpdate(() => {
        const sectionId = $route.params.section;
        console.log($route.params.section);
        const element = document.getElementById(sectionId);
        element.scrollIntoView();
      });
    }
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

      section,
      options
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
    max-width: 400px

</style>
