import axios from "axios";

const URL = process.env.VUE_APP_baseURL;
const baseApiClient = axios.create({
  // baseURL: "http://novotestapi.flattrade.in:29091",
  // baseURL: "https://novoapi.flattrade.in",
  // baseURL: "http://localhost:29091",
  baseURL: URL,
  withCredentials: true, //this is default
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

export default {
  token(input) {
    return baseApiClient.get(
      "/token?token=" + input.token + "&clientId=" + input.clientID
    );
  },
  tokenValidation() {
    return baseApiClient.get("/tokenValidation");
  },
  DeleteCookie() {
    return baseApiClient.get("/deleteCookie");
  },
  // logout
  logOut() {
    return baseApiClient.get("/logout");
  },
  // verify client as admin
  verifyClient() {
    return baseApiClient.get("/verifyClient");
  },
  GetActiveIpo() {
    return baseApiClient.get("/getIpoMaster");
  },
  // GetUpi() {
  //   return baseApiClient.get("/getUpi");
  // },
  GetHistory() {
    return baseApiClient.get("/getIpoHistory");
  },
  GetReport(body) {
    return baseApiClient.post("/getReport", body);
  },
  DefaultReport() {
    return baseApiClient.post("/getDefaultReport");
  },
  FetchManual() {
    return baseApiClient.get("/getManual");
  },
  PlaceOrder(body) {
    return baseApiClient.post("/placeOrder", body);
  },
  Adduser(body) {
    return baseApiClient.post("/addUser", body);
  },
  ChangePassword(body) {
    return baseApiClient.post("/changePassword", body);
  },
  GetModify(input, code) {
    const head = {
      headers: { ID: input, CATEGORY: code },
    };
    return baseApiClient.get("/getModify", head);
  },
  GetHistoryRecors(input, number) {
    const head = {
      headers: { ID: input, NO: number },
    };
    return baseApiClient.get("/getHistoryRecords", head);
  },
  RouterValidation(route) {
    const head = {
      headers: { ROUTER: route },
    };
    return baseApiClient.get("/validRouter", head);
  },
  GetDirectory() {
    return baseApiClient.get("/fetchDirectory");
  },
  // Get Isin and Blog link from API
  GetBlogLink() {
    return baseApiClient.get("/getBlogLink");
  },
  AddBlogLink(body) {
    return baseApiClient.post("/addBlogLink", body);
  },
  GetSgbMaster() {
    return baseApiClient.get("/getSgbMaster");
  },
  // AUG 14 2023
  SetMemberDetail(body) {
    return baseApiClient.post("/setMemberDetail", body);
  },
  GetMemberDetail() {
    return baseApiClient.get("/getMemberDetail");
  },
  GetSubMenu(ParentId) {
    const head = {
      headers: { ID: ParentId },
    };
    return baseApiClient.get("/getSubMenu", head);
  },
  //-----------Aug 17--------
  SGBPlaceOrder(body) {
    return baseApiClient.post("/sgb/placeOrder", body);
  },
  //-----------Aug 18--------
  FetchFund() {
    return baseApiClient.get("/sgb/fetchFund");
  },
  // To get SGB Modify detail
  GetSgbModifyDetail(MasterId, OrderNo) {
    const head = {
      headers: { ID: MasterId, ORDERNO: OrderNo },
    };
    return baseApiClient.get("/sgb/getModify", head);
  },
  // TO get SGB History
  GetSgbHistory() {
    return baseApiClient.get("/sgb/getSgbHistory");
  },
  GetRoleTask() {
    return baseApiClient.get("/GetRoleTask");
  },
  SetRole(Rolemaster) {
    return baseApiClient.put("/setRole", Rolemaster);
  },
  SetTask(Taskmaster) {
    return baseApiClient.put("/setTask", Taskmaster);
  },
  GetRoleTaskMaster() {
    return baseApiClient.get("/GetRoleTaskMaster");
  },
  SetRoleTaskMaster(RoleTask) {
    return baseApiClient.put("/setRoleTaskMaster", RoleTask);
  },
  // TO Get Clients Demat balance
  GetDomainList() {
    return baseApiClient.get("/getBrokerList");
  },
  SetDomain(domain) {
    return baseApiClient.post("/setDomain", domain);
  },
  GetBrokerList() {
    return baseApiClient.get("/getAdminList");
  },
  GetRedirectURL() {
    return baseApiClient.get("/getRedirectUrl");
  },
  // TO Get Clients Name
  GetClientName() {
    return baseApiClient.get("/getClientName");
  },
  GetMemberUser() {
    return baseApiClient.get("/getMemberUser");
  },

  //======================= Dashboard =========================
  GetDashboardDetail(route) {
    const head = {
      headers: { Path: route },
    };
    return baseApiClient.get("/dashboard", head);
  },

  // ========================== LookUp Addtion ================================

  GetLookUpHeader() {
    return baseApiClient.get("/LookUpHeader");
  },
  FetchLookUpDetails(id) {
    return baseApiClient.put("/LookUpDetails", id);
  },
  AddLookUpHeader(input) {
    return baseApiClient.put("/InsertHeader", input);
  },
  AddLookUpDetails(input) {
    return baseApiClient.put("/InsertDetails", input);
  },
  UpdateLookUpHeader(UpdateItem) {
    return baseApiClient.put("/UpdateHeader", UpdateItem);
  },
  UpdateLookUpDetails(input) {
    return baseApiClient.put("/UpdateDetails", input);
  },

  //================================ 10 NOV 2023 =========================
  SetDashbord(body) {
    return baseApiClient.post("/setDashboard", body);
  },
  //================================ 20 NOV 2023 =========================
  ManualOfflineSch() {
    return baseApiClient.get("/manualOfflineSch");
  },

  GetCategory(masterId, path) {
    const head = {
      headers: { ID: masterId, PATH: path },
    };
    return baseApiClient.get("/getCategory", head);
  },
  //================================ 07 DEC 2023 =========================
  GetCategroyPurFlag(masterId) {
    const head = {
      headers: { ID: masterId },
    };
    return baseApiClient.get("/getCategoryPurFlag", head);
  },
  LogDownload() {
    // console.log("LogDownload")
    return baseApiClient.get("/LogDownload");
    //return baseApiClient.get("/LogDownload", { responseType: "blob", });
  },
  GetRegistrarDetails() {
    return baseApiClient.get("/getRegistrarDetails");
  },
  SetRegistrar(registrar) {
    // console.log(registrar)
    return baseApiClient.put("/setRegisterDetails", registrar);
  },
  GetVersion() {
    return baseApiClient.get("/getAllVersions");
  },
  SetVersion(verion) {
    return baseApiClient.put("/setAppVersion", verion);
  },
  // To get current market demand of the active IPO's
  GetIpoMktData(masterId) {
    const head = {
      headers: { ID: masterId },
    };
    return baseApiClient.get("/getIpoMktData", head);
  },
  
  //=========================== NCB ==============================================
  
   // NCB Master
  GetNcbMaster() {
    return baseApiClient.get("/getNcbMaster");
  },

  // History
  getNcbOrderHistory() {
    return baseApiClient.get("/getNcbOrderHistory");
  },

  // PlaceOrder
  NcbPlaceOrder(body) {
    return baseApiClient.post("/ncb/ncbPlaceOrder", body);
  },
  // To get SGB Modify detail
  GetNcbModifyDetail(MasterId, OrderNo) {
    const head = {
      headers: { ID: MasterId, ORDERNO: OrderNo },
    };
    return baseApiClient.get("/ncb/getNcbModify", head);
  },
  GetMasterControl(){
    return baseApiClient.get("/getMasterControl");
  },
  SetMasterControl(item,CurrentTittle){
    const hdr = {
      headers: { CurrentTittle:CurrentTittle  },
    };
    return baseApiClient.post("/setMasterControl", item,hdr);
  },
  GmailReader(Date){
    const head = {
      headers: { setDate: Date },
    };
    return baseApiClient.get("/gmailRetriever", head);
  }

};
