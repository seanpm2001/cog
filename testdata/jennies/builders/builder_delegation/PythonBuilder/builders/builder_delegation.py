import typing
from ..cog import builder as cogbuilder
from ..models import builder_delegation


class DashboardLink(cogbuilder.Builder[builder_delegation.DashboardLink]):    
    __internal: builder_delegation.DashboardLink

    def __init__(self):
        self.__internal = builder_delegation.DashboardLink()

    def build(self) -> builder_delegation.DashboardLink:
        return self.__internal    
    
    def title(self, title: str) -> typing.Self:        
        self.__internal.title = title
    
        return self
    
    def url(self, url: str) -> typing.Self:        
        self.__internal.url = url
    
        return self
    

class Dashboard(cogbuilder.Builder[builder_delegation.Dashboard]):    
    __internal: builder_delegation.Dashboard

    def __init__(self):
        self.__internal = builder_delegation.Dashboard()

    def build(self) -> builder_delegation.Dashboard:
        return self.__internal    
    
    def id_val(self, id_val: int) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def title(self, title: str) -> typing.Self:        
        self.__internal.title = title
    
        return self
    
    def links(self, links: list[cogbuilder.Builder[builder_delegation.DashboardLink]]) -> typing.Self:        
        links_resources = [r.build() for r in links]
        self.__internal.links = links_resources
    
        return self
    
    def single_link(self, single_link: cogbuilder.Builder[builder_delegation.DashboardLink]) -> typing.Self:        
        single_link_resource = single_link.build()
        self.__internal.single_link = single_link_resource
    
        return self
    