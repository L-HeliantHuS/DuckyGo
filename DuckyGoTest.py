import requests

# 全局变量
mainURL = "http://127.0.0.1:8000/api/v1"  # 服务器URL
sess = requests.Session()  # 全局Session会话维持

# 测试的用户名和密码 最终的用户名为 "testUser" 密码为 "newtestpassword."
nickName = "testUser"
userName = "testUser"
password = "testpassword."
newPassword = "newtestpassword."


# Ping 服务器心跳测试
def Ping():
    response = requests.get(f"{mainURL}/ping").json()
    try:
        if response.get("msg") == "Pong":
            print("[+]ping接口没问题~")
    except:
        print("[-]ping接口出现问题.")


# Register注册用户
def Register():
    data = {
        "nickname": nickName,
        "user_name": userName,
        "password": password,
        "password_confirm": password
    }

    response = sess.post(f"{mainURL}/user/register", data=data).json()
    if response.get("code") == 0:
        print("[+]注册接口没问题~")
    else:
        print(f"[-]Register: {response.get('code')} | {response.get('msg')}")


# Login 登录
def Login():
    data = {
        "user_name": userName,
        "password": password
    }
    response = sess.post(f"{mainURL}/user/login", data=data).json()
    if response.get("code") == 0:
        if sess.cookies.get("gin-session") is not None:
            print("[+]登录接口没问题")
        else:
            print("[-]登录接口没问题, 可是没有返回cookie!")
    else:
        print(f"[-]Login:  {response.get('code')} | {response.get('msg')}")


# Me 获取个人信息
def Me():
    response = sess.get(f"{mainURL}/user/me").json()
    if response.get("code") == 0:
        print("[+]个人信息接口没问题~")
    else:
        print(f"[-]Me:  {response.get('code')} | {response.get('msg')}")


# ChangePassword 更改密码
def ChangePassword():
    data = {
        "password": newPassword,
        "password_confirm": newPassword
    }
    response = sess.put(f"{mainURL}/user/changepassword", data=data).json()
    if response.get("code") == 0:
        print(f"[+]更改密码没问题~")
    else:
        print(f"[-]ChangePassword:  {response.get('code')} | {response.get('msg')}")


# Logout 注销
def Logout():
    response = sess.delete(f"{mainURL}/user/logout").json()
    if response.get("code") == 0:
        print("[+]注销接口没问题~")
    else:
        print(f"[-]Logout:  {response.get('code')} | {response.get('msg')}")


if __name__ == "__main__":
    Ping()
    Register()
    Login()
    Me()
    ChangePassword()
    Logout()
