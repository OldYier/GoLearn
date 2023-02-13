package leetcode

type AuthenticationManager struct {
	tokenId     string
	currentTime int8
	timeToLive  int8
}

func Constructor(timeToLive int) AuthenticationManager {
	aut := AuthenticationManager{}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {

}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {

}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {

}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
//  AuthenticationManager authenticationManager = new AuthenticationManager(5); // 构造 AuthenticationManager ，设置 timeToLive = 5 秒。
//  authenticationManager.renew("aaa", 1); // 时刻 1 时，没有验证码的 tokenId 为 "aaa" ，没有验证码被更新。
//  authenticationManager.generate("aaa", 2); // 时刻 2 时，生成一个 tokenId 为 "aaa" 的新验证码。
//  authenticationManager.countUnexpiredTokens(6); // 时刻 6 时，只有 tokenId 为 "aaa" 的验证码未过期，所以返回 1 。
//  authenticationManager.generate("bbb", 7); // 时刻 7 时，生成一个 tokenId 为 "bbb" 的新验证码。
//  authenticationManager.renew("aaa", 8); // tokenId 为 "aaa" 的验证码在时刻 7 过期，且 8 >= 7 ，所以时刻 8 的renew 操作被忽略，没有验证码被更新。
//  authenticationManager.renew("bbb", 10); // tokenId 为 "bbb" 的验证码在时刻 10 没有过期，所以 renew 操作会执行，该 token 将在时刻 15 过期。
//  authenticationManager.countUnexpiredTokens(15); // tokenId 为 "bbb" 的验证码在时刻 15 过期，tokenId 为 "aaa" 的验证码在时刻 7 过期，所有验证码均已过期，所以返回 0 。
