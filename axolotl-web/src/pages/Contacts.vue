<template>
  <div class="contact-list">
    <div v-if="error !== null" v-translate class="alert alert-danger">
      Can't change contact list: {{ error }}
    </div>
    <div v-if="importing" v-translate class="alert alert-warning">
      Importing contacts, head back later
    </div>
    <div v-if="showActions" class="actions-header">
      <button class="btn" @click="delContact(i)">
        <font-awesome-icon icon="trash" />
      </button>
      <button class="btn" @click="editContactModalOpen(contact, i)">
        <font-awesome-icon icon="pencil-alt" />
      </button>
      <button class="btn hide-actions">
        <font-awesome-icon icon="times" @click="closeActionMode" />
      </button>
    </div>
    <div v-if="contacts.length === 0" v-translate class="empty">
      Contact list is empty...
    </div>
    <div v-if="contactsFilterActive">
      <div
        v-for="contact in contactsFiltered"
        :key="contact.Tel"
        :class="
          contact.Tel === editContactId
            ? 'selected btn col-12 chat'
            : 'btn col-12 chat'
        "
      >
        <div class="row chat-entry">
          <div
            :class="'avatar col-3 ' + checkForUUIDClass(contact)"
            @click="contactClick(contact, i)"
          >
            <div class="badge-name">
              {{ contact.Name[0] + contact.Name[1] }}
            </div>
          </div>
          <div
            class="meta col-9"
            data-long-press-delay="500"
            @click="contactClick(contact)"
            @long-press="showContactAction(contact)"
          >
            <p class="name">{{ contact.Name }}</p>
            <p class="number">{{ contact.Tel }}</p>
          </div>
        </div>
      </div>
    </div>
    <div
      v-for="contact in contacts"
      v-else
      :key="contact.Tel"
      :class="
        contact.Tel === editContactId
          ? 'selected btn col-12 chat'
          : 'btn col-12 chat'
      "
    >
      <div class="row chat-entry">
        <div
          :class="'avatar col-3 avatar ' + checkForUUIDClass(contact)"
          @click="contactClick(contact, i)"
        >
          <div class="badge-name">{{ contact.Name[0] + contact.Name[1] }}</div>
        </div>
        <div
          class="meta col-9"
          data-long-press-delay="500"
          @click="contactClick(contact)"
          @long-press="showContactAction(contact)"
        >
          <p class="name">{{ contact.Name }}</p>
          <p class="number">{{ contact.Tel }}</p>
        </div>
      </div>
    </div>

    <div v-if="addContactModal" class="addContactModal">
      <add-contact-modal
        @close="addContactModal = false"
        @add="addContact($event)"
      />
    </div>
    <div v-if="editContactModal" class="editContactModal">
      <edit-contact-modal
        :id="contactId"
        :contact="contact"
        @close="editContactModal = false"
        @save="saveContact($event)"
      />
    </div>
    <div v-if="startChatModal" class="startChatModal">
      <start-chat-modal @close="startChatModal = false" />
    </div>
    <button class="btn add-contact" @click="addContactModal = true">
      <font-awesome-icon icon="plus" />
    </button>
  </div>
</template>

<script>
import AddContactModal from "@/components/AddContactModal.vue";
import EditContactModal from "@/components/EditContactModal.vue";
import StartChatModal from "@/components/StartChatModal.vue";
import { validateUUID } from "@/helpers/uuidCheck";
import longPressEvent from "long-press-event/dist/long-press-event.min.js";

export default {
  name: "Contacts",
  components: {
    AddContactModal,
    EditContactModal,
    StartChatModal,
  },
  data() {
    return {
      addContactModal: false,
      showActions: false,
      editContactModal: false,
      contact: null,
      contactId: null,
      startChatModal: false,
      editContactId: "",
      i: null,
    };
  },
  computed: {
    contacts() {
      return this.$store.state.contacts;
    },
    contactsFiltered() {
      return this.$store.state.contactsFiltered;
    },
    contactsFilterActive() {
      return this.$store.state.contactsFilterActive;
    },
    error() {
      return this.$store.state.rateLimitError;
    },
    importing() {
      return this.$store.state.importingContacts;
    },
  },
  mounted() {
    this.$store.dispatch("getContacts");
  },
  methods: {
    validateUUID,
    addContact(data) {
      this.$store.dispatch("addContact", data);
      this.addContactModal = false;
    },
    delContact() {
      this.$store.dispatch("delContact", this.editContactId);
      this.showActions = false;
      this.editContactId = "";
    },
    checkForUUIDClass(contact) {
      var isValid = this.validateUUID(contact.UUID);
      return isValid ? "" : "not-registered";
    },
    saveContact(data) {
      this.editContactModal = false;
      this.showActions = false;
      this.editContactId = "";
      this.$store.dispatch("editContact", data);
    },
    showContactAction(contact) {
      this.editContactId = contact.Tel;
      this.contact = contact;
      this.showActions = true;
    },
    closeActionMode() {
      this.addContactModal = false;
      this.showActions = false;
      this.editContactModal = false;
      this.contact = null;
      this.contactId = null;
      this.startChatModal = false;
      this.editContactId = "";
    },
    contactClick(contact) {
      if (!this.showActions) {
        if (this.validateUUID(contact.UUID))
          this.$store.dispatch("createChat", contact.UUID);
      } else {
        this.editContactId = contact.Tel;
      }
    },
    editContactModalOpen() {
      this.editContactModal = true;
      this.contactId = this.editContactId;
      this.showActions = false;
      this.editContactId = "";
    },
    startChatModalOpen() {
      if (!this.showActions) {
        this.startChatModal = true;
      }
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.btn.add-contact {
  position: fixed;
  bottom: 16px;
  right: 10px;
  background-color: #2090ea;
  color: #fff;
  border-radius: 50%;
  width: 45px;
  height: 45px;
  font-size: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.chat {
  padding: 0px;
}
.number {
  font-size: 14px;
}
.actions-header {
  position: fixed;
  background-color: #173d5c;
  width: 100%;
  left: 0;
  display: flex;
  justify-content: flex-end;
  z-index: 2;
  top: 0;
  height: 51px;
}
.hide-actions {
  padding-right: 40px;
}
.col-2.actions {
  position: absolute;
  display: flex;
  right: 0px;
  justify-content: center;
  align-items: center;
}
.col-2.actions .btn {
  font-size: 15px;
  padding: 5px;
}
.selected {
  background-color: #c5e4f0;
}
.empty {
  width: 100%;
  height: 70vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
.not-registered .badge-name {
  background: linear-gradient(
    0deg,
    rgb(191, 191, 191) 8%,
    rgb(100, 100, 100) 42%,
    rgb(134, 134, 134) 100%
  );
}
</style>
