package builder_delegation;


public class DashboardLink {
    public String title;
    public String url;
    
    public static class Builder {
        private DashboardLink internal;
        
        public Builder() {
            this.internal = new DashboardLink();
        }
    public Builder setTitle(String title) {
    this.internal.title = title;
        return this;
    }
    
    public Builder setUrl(String url) {
    this.internal.url = url;
        return this;
    }
    public DashboardLink build() {
            return this.internal;
        }
    }
}