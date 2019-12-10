1-Event Driven Cominication yapılması,
    2-Gelen eventlerin parçalanarak ilgili listenerların ilgili handlerlarını tetiklemesi,
3- Event based cominication için modellerin oluşturulması,
    4-Game boardın oluşturulması
    5-Server tarafının mimarisini şekillendirerek ayağa kaldırılması
6-Client tarafının mimarisi ve  servera bağlanabilmesi,
7-Bağlanan clientların sıraya girebilmesini sağlayacak yapının kurulması,
8-Server kendi içerisinde 1 den fazla gameBoard barındırmalıdır.
9-Clientlar bağlandığında username belirleyebilmelidir.
    10-Clientlar tarafından gönderilicek eventlerin belirlenmesi
11-Main metodunda GameEvent sıralamasının yapılması,
12-Frame rate belirlenecek döngü hep kendini takip edecek,
13-Game loop dışında eventler modelleri değiştirebilmeli,
14-Clientlar ile pingleşilerek kimin requestlere cevap verip vermediği anlaşılmalı.

----
(10)=> clientlar tarafından gönderilip dinlenecek eventler.
    a-Top sopaya çarptı
    b-Sopaların boostu bastı
    c-sopa hareket etti(sağa git sola git dur)
    d-gol oldu (golü atan ve yiyen state tutuluyor)

(4)=> Game board 
    game array=>
        x,y arrayLength
        sopaların kendi alanları (client ID'lerine göre belirlenebilir)[]sopalar
        top
        
    top=>
        x,y vector,
        hız,
        topa son vuran,
    sopalar=>
        x,y coordinanat,
        yön,
        hız,
        clientId,
        Renk,
        bölge
(2)=>EventListener
    UserConnected,
    UserJoined,
    UserJoinedQueue,
    UserLeft,
    GameStart,
    SopaMoved,
    TopSopaya çarptı,
    ScorEldeEdildi,

(5)=> Server struct
        dispatcher,
        clients,
        pattern,
        logger,
        idManager,
        userNameRegistry,
        upgrader,
    =>NewServer(pattern,logger,dispatcher,idManager,USerNameRegistry)

    =>ServerListen()
        handlerFunc=>fireUserConnected,logger,healtcheckHandler
    =>SendToClient(clientId,meesage) =>send clientMessage,
    =>GetClientIDs,
    =>GetClient,

(6)=>Client struct
            id,
            *websocket,
            channel,
            doneCh,
            server *Server,
    NewClient(ws,server)=>return make2Chan,idManager.NexTPlayerId(),ws
    SendMessage([]bytes)=>channel<-bytes=>logger,
    ClientListen=>listenWrite(),listenRead()
        listenWrite()=>while wait channels=>ws.WriteMessage,
        listenRead()=>while wait ReadWebSocket=>Handle EvensMessage And Fire EventHandler(switchCase,FireTargetAngle,FireJoinedGame etc.)





        