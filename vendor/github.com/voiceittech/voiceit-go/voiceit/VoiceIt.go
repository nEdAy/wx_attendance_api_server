package voiceit
import "fmt"
import "net/http"
import "crypto/sha256"
import "io"
import "encoding/hex"
import "io/ioutil"
import "bytes"

type VoiceIt struct {
    devID string
}

func New(developerID string) *VoiceIt{
    return &VoiceIt{
        devID: developerID,
    }
}

func (v *VoiceIt) CreateUser(userId string, passwd string) string{
    hasher := sha256.New()
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/users", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("UserId" , userId)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) GetUser(userId string, passwd string) string{
    hasher := sha256.New()
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("GET", "https://siv.voiceprintportal.com/sivservice/api/users", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("UserId" , userId)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) DeleteUser(userId string, passwd string) string{
    hasher := sha256.New()
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("DELETE", "https://siv.voiceprintportal.com/sivservice/api/users", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("UserId" , userId)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result

}

func (v *VoiceIt) CreateEnrollment(userId string, passwd string, pathToEnrollmentWav string, contentLanguage ... string) string {
    hasher := sha256.New()
    contentLang :=""
    if len(contentLanguage) > 0 {
        contentLang = contentLanguage[0]
    } else {
        contentLang =""
    }
    wavData, err := ioutil.ReadFile(pathToEnrollmentWav)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/enrollments", bytes.NewReader(wavData))
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("UserId" , userId)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("ContentLanguage" , contentLang)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err:= ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) CreateEnrollmentByWavURL(userId string, passwd string, urlToEnrollmentWav string, contentLanguage ... string) string {
    hasher := sha256.New()
    contentLang :=""
    if len(contentLanguage) > 0 {
        contentLang = contentLanguage[0]
    } else {
        contentLang =""
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/enrollments/bywavurl", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("UserId" , userId)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("VsitwavURL", urlToEnrollmentWav)
    req.Header.Add("ContentLanguage" , contentLang)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err:= ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) GetEnrollments(userId string, passwd string) string{
    hasher := sha256.New()
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("GET", "https://siv.voiceprintportal.com/sivservice/api/enrollments", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("UserId" , userId)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) DeleteEnrollment(userId string, passwd string, enrollmentId string) string{
    hasher := sha256.New()
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("DELETE", "https://siv.voiceprintportal.com/sivservice/api/enrollments/"+enrollmentId, nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("UserId" , userId)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result

}

func (v *VoiceIt) Authentication(userId string, passwd string, pathToAuthenticationWav string,contentLanguage ... string) string {
    hasher := sha256.New()
    contentLang :=""
    if len(contentLanguage) > 0 {
        contentLang = contentLanguage[0]
    } else {
        contentLang =""
    }
    wavData, err := ioutil.ReadFile(pathToAuthenticationWav)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/authentications", bytes.NewReader(wavData))
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("UserId" , userId)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("ContentLanguage" , contentLang)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err:= ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) AuthenticationByWavURL(userId string, passwd string, urlToAuthenticationWav string, contentLanguage ... string) string {
    hasher := sha256.New()
    contentLang :=""
    if len(contentLanguage) > 0 {
        contentLang = contentLanguage[0]
    } else {
        contentLang =""
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/authentications/bywavurl", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("UserId" , userId)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("VsitwavURL", urlToAuthenticationWav)
    req.Header.Add("ContentLanguage" , contentLang)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err:= ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}
