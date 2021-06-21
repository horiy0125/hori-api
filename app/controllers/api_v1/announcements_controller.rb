class ApiV1::AnnouncementsController < ApplicationController

  def index
    @announcements = Announcement.order(created_at: :desc)
    render :json => @announcements
  end
end
