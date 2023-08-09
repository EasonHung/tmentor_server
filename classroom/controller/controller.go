package controller

// "mentor/classroom/service/class_process_service"

// var upGrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func Route(route *gin.Engine) {
// 	classroomRoute := route.Group("/websocket")
// 	{
// 		classroomRoute.GET("/", classroomWs)
// 		classroomRoute.GET("/classroomList", getClassroomMap)
// 		classroomRoute.GET("/redisChannel/list", getRedisChannelList)
// 	}
// }

// func classroomWs(c *gin.Context) {
// 	var (
// 		conn *websocket.Conn
// 		err  error
// 	)
// 	// upgrade to websocket connection
// 	if conn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
// 		return
// 	}

// 	var applicationType string
// 	applicationTypeHeader, _ := c.GetQuery("applicationType")
// 	if applicationTypeHeader == "web" {
// 		applicationType = class_channel_var.WEB
// 	} else {
// 		applicationType = class_channel_var.APP
// 	}

// 	userId, _ := c.GetQuery("userId")
// 	defer func(userId string, applicationType string) {
// 		err := class_channel_service.NotifyClosureAndCleanMap(userId, applicationType)
// 		if err != nil {
// 			log.Logger.Errorf("%+v\n", err)
// 		}
// 		log.Logger.Info("close connect " + userId)
// 		conn.Close()
// 	}(userId, applicationTypeHeader)

// 	// add wesocket connection node
// 	node := &dto.WebSocketNode{
// 		Connection: conn,
// 		DataQueue:  make(chan []byte, 50),
// 	}
// 	reflect.ValueOf(node)

// 	// 訂閱
// 	sub := mentor_redis.Client.Subscribe(class_channel_var.WS_REDIS_PREFIX + applicationType + userId)
// 	defer func(userId string, sub *redis.PubSub) {
// 		conn.Close()
// 		sub.Unsubscribe(class_channel_var.WS_REDIS_PREFIX + applicationType + userId)
// 		sub.Close()
// 		log.Logger.Info("dispatch")
// 	}(userId, sub)

// 	go listenAndSendToClient(sub, node)
// 	recieveAndPublish(node)
// 	log.Logger.Info("disconnect")
// }

// func listenAndSendToClient(sub *redis.PubSub, node *dto.WebSocketNode) {
// 	for msg := range sub.Channel() {
// 		err := node.Connection.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
// 		if err != nil {
// 			log.Logger.Error("[system] error Occured during sending single chat message")
// 			return
// 		}
// 	}
// 	log.Logger.Info("stop listening")
// }

// func recieveAndPublish(senderNode *dto.WebSocketNode) {
// 	for {
// 		_, data, err := senderNode.Connection.ReadMessage()
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			log.Logger.Error("[system] error Occured during recieving single chat message", err)
// 			return
// 		}

// 		// json unmarshal
// 		message := dto.Message{}
// 		if err := json.Unmarshal(data, &message); err != nil {
// 			log.Logger.Errorf("[system] error Occured during unmarshal message %+v", err)
// 			return
// 		}

// 		switch message.Cmd {
// 		// // todo: close room
// 		// case ws_cmd.OpenRoom:
// 		// 	// todo: verify user token
// 		// 	openClassroom(message)
// 		// case ws_cmd.Ask:
// 		// 	class_channel_service.AskAccess(&message)
// 		// case ws_cmd.Reject:
// 		// 	class_channel_service.RejectAccess(message.RecieverId, message.ReceiverApplicationType, data)
// 		// case ws_cmd.Accept:
// 		// 	acceptAndSendClassroomToken(&message)
// 		// case ws_cmd.JoinRoom:
// 		// 	class_channel_service.JoinClassRoom(&message)
// 		// case ws_cmd.InstantMessage:
// 		// 	class_channel_service.SendInstantMsg(&message)

// 		// start class procedure
// 		// case ws_cmd.LastClassInfo:
// 		// 	class_channel_service.GetLastClassInfo(message)
// 		// case ws_cmd.InitClass:
// 		// 	class_channel_service.InitClass(message)
// 		// case ws_cmd.AcceptClass:
// 		// 	class_channel_service.AcceptClass(message)
// 		// case ws_cmd.FinishClass:
// 		// 	class_channel_service.FinishClass(message)

// 		// webRTC
// 		// case ws_cmd.Offer:
// 		// 	class_channel_service.SendOffer(message.ClassroomId, message.SenderId, message.ApplicationType, data)
// 		// case ws_cmd.Answer:
// 		// 	class_channel_service.SendAnswer(message.ClassroomId, message.SenderId, message.ApplicationType, data)
// 		// case ws_cmd.Candidate:
// 		// 	class_channel_service.SendCandidate(message.ClassroomId, message.SenderId, message.ApplicationType, data)
// 		// case ws_cmd.Hangup:
// 		// 	class_channel_service.SendHangUp(message.ClassroomId, message.SenderId, message.ApplicationType, data)
// 		default:
// 			log.Logger.Warn("[system] wrong cmd", message.Cmd)
// 		}
// 	}
// }

// func getRedisChannelList(c *gin.Context) {
// 	// 没有指定查询channel的匹配模式，则返回所有的channel
// 	// 匹配user_开头的channel
// 	chs, _ := mentor_redis.Client.PubSubChannels(class_channel_var.WS_REDIS_PREFIX + "*").Result()
// 	res := make([]string, 0)
// 	for _, ch := range chs {
// 		res = append(res, ch)
// 	}

// 	c.JSON(200, res)
// 	return
// }

// func getClassroomMap(c *gin.Context) {
// 	c.JSON(200, class_channel_service.GetClassroomList())
// }

// func getClassroomToken(c *gin.Context) {
// 	mentorId, _ := c.Get("userId")

// 	err, token := classroom_registry_service.GetUserClassroomToken(mentorId.(string))
// 	if err != nil {
// 		log.Logger.Errorf("%+v\n", err)
// 		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
// 		return
// 	}

// 	c.JSON(200, res.NewSuccessResponse(token))
// 	return
// }

// // func openClassroom(senderMessage dto.Message) {
// // 	err := class_process_service.OpenClassroom(senderMessage)
// // 	if err != nil {
// // 		senderMessage.ToErrorCmd(err.Error())
// // 		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
// // 		return
// // 	}
// // }

// func acceptAndSendClassroomToken(dto *dto.Message) {
// 	err, classroomToken := classroom_registry_service.GetUserClassroomToken(dto.SenderId)
// 	if err != nil {
// 		log.Logger.Errorf("%+v\n", err)
// 		return
// 	}

// 	dto.Message = classroomToken
// 	data, _ := json.Marshal(dto)
// 	mentor_redis.Client.Publish(class_channel_var.WS_REDIS_PREFIX+dto.ReceiverApplicationType+"-"+dto.RecieverId, string(data))
// }
