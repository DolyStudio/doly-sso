# Doly-SSO

## Todo

- [x] Google Social Login
- [ ] Naver Social Login
- [ ] Kakao Social Login

## swagger 등록하는 법

```sh
    ## in repository
    go install github.com/swaggo/swag/cmd/swag@latest

    ## swag error
    zsh: command not found: swag

    ## swag error solution
    export PATH=$(go env GOPATH)/bin:$PATH

    ## fmt (formatting)
    swag fmt

    ## init (init)
    swag init

    ## import 등록

    import (
        ...
        _ "{application}/docs
    )
```

## Google Login Logic

```js
async function handleGoogleLogin() {
    const button = document.querySelector('.social-button.google');

    try {
        setLoadingState(button, true);
        showResult('구글 로그인 요청 중...', false);

        const response = await fetch(`${BACKEND_URL}/auth/google/login`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }
        });

        if (response.ok) {
            const data = await response.json();

            // Google 로그인 페이지로 리다이렉트
            window.location.href = data.auth_url;

        } else {
            const errorData = await response.json();
            showResult(`구글 로그인 실패!\n에러: ${JSON.stringify(errorData, null, 2)}`, false);
        }
    } catch (error) {
        showResult(`구글 로그인 오류!\n에러: ${error.message}`, false);
    } finally {
        setLoadingState(button, false);
    }
}
```
