import requests

# 全局变量
mainUrl = "http://127.0.0.1:8000"
V1URL = f"{mainUrl}/api/v1"  # v1服务器URL
V2URL = f"{mainUrl}/api/v2"  # v2服务器URL
sess = requests.Session()  # 全局Session会话维持

# 测试的用户名和密码 最终的用户名为 "testUser" 密码为 "newtestpassword."
nickName = "testUser"
userName = "testUser"
password = "testpassword."
newPassword = "newtestpassword."


# Ping 服务器心跳测试
def Ping():
    try:
        response = requests.get(f"{mainUrl}/ping").json()
        if response.get("msg") == "Pong":
            return True
    except:
        return False


class V1Test:
    # Register注册用户
    @staticmethod
    def Register():
        data = {
            "nickname": nickName,
            "user_name": userName,
            "password": password,
            "password_confirm": password
        }

        response = sess.post(f"{V1URL}/user/register", data=data).json()
        if response.get("code") == 0:
            print("[+]注册接口没问题~")
        else:
            print(f"[-]Register: {response.get('code')} | {response.get('msg')}")


    # Login 登录
    @staticmethod
    def Login():
        data = {
            "user_name": userName,
            "password": password
        }
        response = sess.post(f"{V1URL}/user/login", data=data).json()
        if response.get("code") == 0:
            if sess.cookies.get("gin-session") is not None:
                print("[+]登录接口没问题")
            else:
                print("[-]登录接口没问题, 可是没有返回cookie!")
        else:
            print(f"[-]Login:  {response.get('code')} | {response.get('msg')}")


    # Me 获取个人信息
    @staticmethod
    def Me():
        response = sess.get(f"{V1URL}/user/me").json()
        if response.get("code") == 0:
            print("[+]个人信息接口没问题~")
        else:
            print(f"[-]Me:  {response.get('code')} | {response.get('msg')}")


    # ChangePassword 更改密码
    @staticmethod
    def ChangePassword():
        data = {
            "password": newPassword,
            "password_confirm": newPassword
        }
        response = sess.put(f"{V1URL}/user/changepassword", data=data).json()
        if response.get("code") == 0:
            print(f"[+]更改密码没问题~")
        else:
            print(f"[-]ChangePassword:  {response.get('code')} | {response.get('msg')}")


    # Logout 注销
    @staticmethod
    def Logout():
        response = sess.delete(f"{V1URL}/user/logout").json()
        if response.get("code") == 0:
            print("[+]注销接口没问题~")
        else:
            print(f"[-]Logout:  {response.get('code')} | {response.get('msg')}")


class V2Test:
    def __init__(self):
        self.headers = {
            "Authorization": ""
        }

    def Register(self):
        data = {
            "nickname": nickName,
            "user_name": userName,
            "password": password,
            "password_confirm": password
        }

        response = requests.post(f"{V2URL}/user/register", data=data).json()
        if response.get("code") == 0:
            print("[+]注册接口没问题~")
        else:
            print(f"[-]Register: {response.get('code')} | {response.get('msg')}")

    def Login(self):
        data = {
            "user_name": userName,
            "password": password
        }
        response = requests.post(f"{V2URL}/user/login", data=data).json()
        if response.get("code") == 0:
            if response["data"] != "":
                self.headers["Authorization"] = "Bearer" + " " + response["data"]
                print(f"[+]登录接口没问题~")
            else:
                print(f"[-]Login: 接口没问题, 但是没有获得Token")
        else:
            print(f"[-]Login:  {response.get('code')} | {response.get('msg')}")

    def Me(self):
        response = requests.get(f"{V2URL}/user/me", headers=self.headers).json()
        if response["code"] == 0:
            print("[+]个人信息接口没问题~")
        else:
            print(f"[-]Me: {response.get('code')} | {response.get('msg')}")


    def ChangePassword(self):
        data = {
            "password": newPassword,
            "password_confirm": newPassword
        }
        response = requests.put(f"{V2URL}/user/changepassword", headers=self.headers, data=data).json()
        if response.get("code") == 0:
            print(f"[+]更改密码没问题~")
        else:
            print(f"[-]ChangePassword:  {response.get('code')} | {response.get('msg')}")

    def Logout(self):
        response = requests.delete(f"{V2URL}/user/logout", headers=self.headers).json()
        if response["code"] == 0:
            response = requests.delete(f"{V2URL}/user/logout", headers=self.headers).json()
            if response["code"] == 40002:
                print(f"[+]用户成功注销, 并且已经生效")
            else:
                print(f"[-]Logout: 接口没问题, 但是注销失败")
        else:
            print(f"[-]Logout:  {response.get('code')} | {response.get('msg')}")

if __name__ == "__main__":
    selectServer = input("Select Test Server Name('v1' or 'v2'):")
    if selectServer == 'v1':
        if Ping() == False:
            print("[x] DuckyGo可能不在线，请检查URL配置或者DuckyGo是否开启")
            exit()

        try:
            V1Test.Register()
            V1Test.Login()
            V1Test.Me()
            V1Test.ChangePassword()
            V1Test.Logout()
        except:
            print("[x] DuckyGo可能未开启v1的api组")
    elif selectServer == 'v2':
        if Ping() == False:
            print("[x] DuckyGo可能不在线，请检查URL配置或者DuckyGo是否开启")
            exit()

        try:
            v2 = V2Test()
            v2.Register()
            v2.Login()
            v2.Me()
            v2.ChangePassword()
            v2.Logout()
        except:
            print("[x] DuckyGo可能未开启v2的api组")
    else:
        print("You select wrong.")