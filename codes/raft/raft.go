package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isleader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//
/*
节点三个状态: Follower, Candidate, Leader. 其中Candidate是Follower向Leader转变的一个中间状态
*/

/*
初始: 所有节点处于Follower状态，并各自开启一个随机的倒计时.
当倒计时结束时向其余节点拉票，若获得了大于一半的票数则选举为Leader
*/

/*
Leader election
Fllower会有一个选举倒计时，选举倒计时是在150ms到300ms之间的随机的值,
当倒计时结束的时候，会成为candidate, 自动开始拉票进行新一轮的选举(先投自己一票),
然后发送Request Vote 消息给其他节点， 其余节点如果在此次选举还未投票则将票投给此candidate并重置倒计时
若candidate获取了超过半数的投票，则成为leader。

leader 会定时发送AppendEntries(心跳包)消息给Followers, 此时间间隔叫作heartbeat timeout
Follower 在收到AppendEntries后会返回消息。

此次选举将会持续到直到某一个Follwer停止接受心跳包, 并成为一个Candidate

若有两个节点同时成为Candidate, 并且都只能获得半数投票，那么倒计时重启再次进行选举
*/

/*
log replication
Leader对Follower有heartbeat机制，用AppendEntries消息来完成
client发送消息给leader后， leader会在下一个心跳包将消息发送给follower，
follower收到后会给leader返回，收到超过一般follower返回，leader进入commit状态，并返回消息给client
leader再次发送消息给follower，让follower进入commit状态
*/
import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"../labrpc"
)

// import "bytes"
// import "../labgob"

//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make(). set
// CommandValid to true to indicate that the ApplyMsg contains a newly
// committed log entry.
//
// in Lab 3 you'll want to send other kinds of messages (e.g.,
// snapshots) on the applyCh; at that point you can add fields to
// ApplyMsg, but set CommandValid to false for these other uses.
//
type ApplyMsg struct {
	CommandValid bool
	Command      interface{}
	CommandIndex int
}

type LogInfo struct {
	Command interface{}
	Term    int32
}

//
// A Go object implementing a single Raft peer.
//
type Raft struct {
	mu        sync.Mutex          // Lock to protect shared access to this peer's state
	peers     []*labrpc.ClientEnd // RPC end points of all peers
	persister *Persister          // Object to hold this peer's persisted state
	me        int                 // this peer's index into peers[]
	dead      int32               // set by Kill()

	// Your data here (2A, 2B, 2C).
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.
	timer       *time.Timer
	currentTerm int32
	voteFor     int32
	leaderId    int32
	commitIndex int32
	lastApplied int32
	log         []LogInfo
	nextIndex   []int32
	matchIndex  []int32
	peerCommit  []int32
}

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {

	var term int
	var isleader bool
	// Your code here (2A).
	return term, isleader
}

//
// save Raft's persistent state to stable storage,
// where it can later be retrieved after a crash and restart.
// see paper's Figure 2 for a description of what should be persistent.
//
func (rf *Raft) persist() {
	// Your code here (2C).
	// Example:
	// w := new(bytes.Buffer)
	// e := labgob.NewEncoder(w)
	// e.Encode(rf.xxx)
	// e.Encode(rf.yyy)
	// data := w.Bytes()
	// rf.persister.SaveRaftState(data)
}

//
// restore previously persisted state.
//
func (rf *Raft) readPersist(data []byte) {
	if data == nil || len(data) < 1 { // bootstrap without any state?
		return
	}
	// Your code here (2C).
	// Example:
	// r := bytes.NewBuffer(data)
	// d := labgob.NewDecoder(r)
	// var xxx
	// var yyy
	// if d.Decode(&xxx) != nil ||
	//    d.Decode(&yyy) != nil {
	//   error...
	// } else {
	//   rf.xxx = xxx
	//   rf.yyy = yyy
	// }
}

//
// example RequestVote RPC arguments structure.
// field names must start with capital letters!
//
type RequestVoteArgs struct {
	// Your data here (2A, 2B).
	term         int32
	candidateId  int32
	lastLogIndex int32
	lastLogTerm  int32
}

//
// example RequestVote RPC reply structure.
// field names must start with capital letters!
//
type RequestVoteReply struct {
	// Your data here (2A).
	term        int32
	commitIndex int32
	voteGranted bool
}

type AppendEntriesArgs struct {
	term         int32
	leaderId     int32
	prevLogIndex int32
	prevLogTerm  int32
	entries      []LogInfo
	leaderCommit int32
}

type AppendEntriesReply struct {
	term    int32
	success bool
}

func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {
	if args.leaderCommit == rf.lastApplied {
		// 进入commit状态, 将操作持久化
		reply.success = true
		rf.currentTerm = args.term
		return
	}

	if args.term <= rf.currentTerm {
		reply.success = false
		reply.term = rf.currentTerm
		return
	}

	//执行replicate
	for i := 0; i < len(args.entries); i++ {
		args.entries[i].Command
	}

	reply.success = true
	reply.term = rf.currentTerm
}

//
// example RequestVote RPC handler.
//
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	// Your code here (2A, 2B).
	if rf.currentTerm < args.term && rf.commitIndex <= args.lastLogIndex {
		rf.currentTerm = args.term
		reply.voteGranted = true
		reply.term = args.term
		rf.leaderId = args.candidateId
	} else {
		reply.voteGranted = false
	}
}

//
// example code to send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *reply with RPC reply, so caller should
// pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a reply. If a reply arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost reply.
//
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return.  Thus there
// is no need to implement your own timeouts around Call().
//
// look at the comments in ../labrpc/labrpc.go for more details.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the reply struct with &, not
// the struct itself.
//
func (rf *Raft) sendRequestVote(server int, args *RequestVoteArgs, reply *RequestVoteReply) bool {
	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
	return ok
}

//
// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election. even if the Raft instance has been killed,
// this function should return gracefully.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
//
func (rf *Raft) Start(command interface{}) (int, int, bool) {
	index := -1
	term := -1
	isLeader := true

	// Your code here (2B).
	isLeader = (rf.leaderId == int32(rf.me))
	term = int(rf.currentTerm)
	index = int(rf.commitIndex)
	return index, term, isLeader
}

//
// the tester doesn't halt goroutines created by Raft after each test,
// but it does call the Kill() method. your code can use killed() to
// check whether Kill() has been called. the use of atomic avoids the
// need for a lock.
//
// the issue is that long-running goroutines use memory and may chew
// up CPU time, perhaps causing later tests to fail and generating
// confusing debug output. any goroutine with a long-running loop
// should call killed() to check whether it should stop.
//
func (rf *Raft) Kill() {
	atomic.StoreInt32(&rf.dead, 1)
	// Your code here, if desired.
}

func (rf *Raft) killed() bool {
	z := atomic.LoadInt32(&rf.dead)
	return z == 1
}

//
// the service or tester wants to create a Raft server. the ports
// of all the Raft servers (including this one) are in peers[]. this
// server's port is peers[me]. all the servers' peers[] arrays
// have the same order. persister is a place for this server to
// save its persistent state, and also initially holds the most
// recent saved state, if any. applyCh is a channel on which the
// tester or service expects Raft to send ApplyMsg messages.
// Make() must return quickly, so it should start goroutines
// for any long-running work.
//
func Make(peers []*labrpc.ClientEnd, me int,
	persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me

	// Your initialization code here (2A, 2B, 2C).
	rf.currentTerm = 0
	rf.commitIndex = 0
	rf.lastApplied = 0
	rf.leaderId = -1
	rf.log = make([]LogInfo, 10)
	rf.matchIndex = make([]int32, 10)
	rf.nextIndex = make([]int32, 10)
	rf.peerCommit = make([]int32, len(peers))

	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			rf.timer = time.NewTimer(time.Millisecond*rand.Intn(150) + 150)
			select {
			case <-rf.timer.C:
				//开启新一轮选举
				count := 1
				arg := &RequestVoteArgs{}
				arg.candidateId = int32(rf.me)
				arg.lastLogTerm = rf.currentTerm
				arg.term = rf.currentTerm + 1
				arg.lastLogTerm = rf.commitIndex
				reply := &RequestVoteReply{}
				rf.currentTerm++

				//拉票
				l := len(rf.peers)
				for i := 0; i < l; i++ {
					if i == me {
						continue
					}

					rf.sendRequestVote(i, arg, reply)
					if reply.voteGranted {
						count++
						//设置peer的commitindex
						rf.peerCommit[i] = reply.commitIndex
					}
				}

				if count <= l/2 {
					// 选举票数未超半数，重新计时
					break
				}

				//成为新节点，开始append msg
				rf.leaderId = int32(rf.me)
				//首先把自己和follower进行同步
				//TODO
			}
		}
	}()
	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())

	return rf
}
