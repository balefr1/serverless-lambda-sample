import React, {Component} from 'react'
import {Table,Label,Input,FormGroup,Button, Modal, ModalHeader, ModalBody, ModalFooter} from 'reactstrap'
import axios from 'axios'

class UserAttachmentPage extends Component{
    HOST= "api.example.com";
    API_URL = `https://${this.HOST}/dev/files`;
    DOWNLOAD_URL= `https://${this.HOST}/dev/file/`;

    state = {
        attachments: [],
        newAttachmentModal:false,
        downloadAttachmentModal:false,
        downloadFile: {}
      }
    componentWillMount() {
        this._RefreshAttachments();
      }
    toggleNewAttachmentModal(){
        this.setState({
            newAttachmentModal:!this.state.newAttachmentModal
        });
    }

    toggleDownloadAttachmentModal(){
      this.setState({
          downloadAttachmentModal:!this.state.downloadAttachmentModal
      });
  }
      _RefreshAttachments(){
        axios.get(this.API_URL).then((response)=>{
          this.setState({
            attachments:response.data
          })
        }).catch( (error) =>{
            console.log(error.response)
            alert("Error - " + error)
          });
      }

      AddAttachment() {
        var formData = new FormData();
        var imagefile = document.querySelector('#file');
        formData.append("file", imagefile.files[0]);
        axios.post(this.API_URL,formData,{
            headers: {
                'Content-Type': 'multipart/form-data'
              }
        }).then((response)=>{
          console.log(response);
          this._RefreshAttachments();
          this.setState({
              newAttachmentModal:false
          })
        }).catch( (error) =>{
            console.log(error.response)
            alert("Error - " + error.response.data.error)
          });
      }
      GetAttachment(id){
        axios.post(this.DOWNLOAD_URL+id+"/download").then((response=>{
          console.log(response.data)
          this.state.downloadFile.filename=response.data.name;
          this.state.downloadFile.url=response.data.download_url;
          this.state.downloadFile.size=response.data.size;
          this.toggleDownloadAttachmentModal();

        })).catch( (error) =>{
            console.log(error.response)
            alert("Error - " + error.response.data.error)
          })
      }
      render(){
        let attachments =this.state.attachments.map((attachment) =>{
          return(
            <tr key={attachment.filename}>
              <td>{attachment.filename}</td>
              <td>{attachment.at}</td>
              <td>{attachment.ip}</td>
              <td> <Button color="primary" onClick={this.GetAttachment.bind(this,attachment.filename)}>DOWNLOAD</Button></td>
            </tr>
          )
        });
        return (
          <div className="App" >
            <div>
            <img src={"./sorint-logo.jpeg"} width={350} height={100} />
              <Button color="primary" onClick={this.toggleNewAttachmentModal.bind(this)}>Upload file</Button>
            </div>
          <br></br>
          <h5>List of files:</h5>
          <Modal isOpen={this.state.newAttachmentModal} toggle={this.toggleNewAttachmentModal.bind(this)}>
            <ModalHeader toggle={this.toggleNewAttachmentModal.bind(this)}>Upload new file</ModalHeader>
            <ModalBody>
             <FormGroup>
              <Label for="name">Choose File</Label>
              <Input type="file" name="file" id="file" ></Input>
             </FormGroup>
            </ModalBody>
            <ModalFooter>
              <Button color="primary" onClick={this.AddAttachment.bind(this)}>Upload File</Button>{' '}
              <Button color="secondary" onClick={this.toggleNewAttachmentModal.bind(this)}>Cancel</Button>
            </ModalFooter>
          </Modal>
          <Modal isOpen={this.state.downloadAttachmentModal} toggle={this.toggleDownloadAttachmentModal.bind(this)}>
            <ModalHeader toggle={this.toggleDownloadAttachmentModal.bind(this)}>Download file</ModalHeader>
            <ModalBody>
              Filename: <h5>{this.state.downloadFile.filename}</h5>
              Size: {this.state.downloadFile.size ? this.state.downloadFile.size : "N/A" }
              <br></br>
              <a href={this.state.downloadFile.url} >Download</a>
            </ModalBody>
            <ModalFooter>
              <Button color="secondary" onClick={this.toggleDownloadAttachmentModal.bind(this)}>Cancel</Button>
            </ModalFooter>
          </Modal>
            <Table style={{ background: 'orange' }}>
              <thead>
                <tr>
                  <th>Filename</th>
                  <th>Last Updated</th>
                  <th>Uploaded from</th>
                </tr>
              </thead>
              <tbody>
              {attachments.length ? attachments : <div><tr><td colspan={4}>No Attachments found!</td></tr></div>}
              </tbody>
            </Table>
          </div>
        );
      }
}

export default UserAttachmentPage;

