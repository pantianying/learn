import com.alibaba.fastjson.JSONObject;

public class Start {
    public static void main(String args[]) {
    String jsonStr =
        "{\"type\":\"com.tuya.basic.mq.domain.KafkaMqData\",\"count\":2,\"ct\":1581645199194,\"data\":{\"@type\":\"com.alibaba.fastjson.JSONObject\",\"devId\":\"6cb74d14b645dd0194wxg0\",\"sub\":\"1\",\"gmtModified\":1581645198L,\"activeTime\":1581645198L,\"icon\":\"smart/program_category_icon/kg.png\",\"pid\":\"gjackgkk\",\"source\":\"apollo\",\"ownerId\":\"11377020\",\"gmtCreate\":1581645198L,\"uuid\":\"6cb74d14b645dd0194wxg0\",\"meshId\":\"6c5a599127aa0129dc62ev\",\"mac\":\"086bd7fffed0811a\",\"uid\":\"ay1580706778130rOY6C\",\"name\":\"强电四路开关 3\",\"gwId\":\"6cb74d14b645dd0194wxg0\",\"lang\":\"zh\"},\"groupId\":\"hermes\",\"offsets\":[75239L,75240L],\"rt\":42}";
        JSONObject jsonObject = JSONObject.parseObject(jsonStr);
        System.out.println(jsonObject.toJSONString());
    }
}
