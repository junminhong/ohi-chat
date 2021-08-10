## What is this
ohi-chat backend server

## Feature
- 基本聊天室的CRUD
- 加入聊天室、離開聊天室

## API DOC and Feature
### Chat room
可以善用MongoDB document storage的特性，一個聊天室等於一個文件的概念做設計。
> DB Type: MongoDB
> DB Table: room_info
> DB Column: id, name, limit_amount, privacy, tags(暫時不需要用到)
- [ ] list chat room
> API [GET]: api/v1/chat-room/list
> token: requirement
> request: {}
> response: {result: 200 or 404, data: room_list[]}
- [ ] create chat room
> API [POST]: api/v1/chat-room/create
> token: requirement
> request: {room_id, room_name, room_limit_amount, room_privacy}
> response: {result: 200 or 404, data: null, message: create chat room successful or failed}
- [ ] edit chat room info
> API [PUT]: api/v1/chat-room/edit
> token: requirement
> request: {room_id, room_name, room_limit_amount, room_privacy}
> response: {result: 200 or 404, data: null, message: edit chat room successful or failed}
- [ ] delete chat room
> API [DELETE]: api/v1/chat-room/delete
> token: requirement
> request: {room_id}
> response: {result: 200 or 404, data: null, message: edit chat room successful or failed}
- [ ] join chat room
> API [POST]: api/v1/chat-room/join
> token: requirement
> request: {room_id}
> response: {result: 200 or 404, data: null, message: join chat room successful or failed}
- [ ] leave chat room
> API [POST]: api/v1/chat-room/leave
> token: requirement
> request: {room_id}
> response: {result: 200 or 404, data: null, message: leave chat room successful or failed}