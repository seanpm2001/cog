package sandbox;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class SomeStruct { 
    @JsonProperty("time")
    public Object time;
    
    public String ToJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements cog.Builder<SomeStruct> {
        private SomeStruct internal;
        
        public Builder() {
            this.internal = new SomeStruct();
        }
    public Builder Time(String from,String to) {
		if (this.internal.time == null) {
			this.internal.time = new Object();
		}
    this.internal.time.from = from;
    this.internal.time.to = to;
        return this;
    }
    public SomeStruct Build() {
            return this.internal;
        }
    }
}
