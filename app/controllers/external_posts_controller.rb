require 'net/http'

class ExternalPostsController < ApplicationController

  before_action :find_external_post, only: %i(edit update destroy)

  DEPLOY_HOOKS_URI = URI.parse(ENV['PORTFOLIO_FRONT_DEPLOY_HOOKS_URL'])

  def index
    @external_posts = ExternalPost.order(created_at: :desc)
  end

  def new
    @external_post = ExternalPost.new
  end

  def create
    ExternalPost.create!(allowed_params)
    if Rails.env.production?
      Net::HTTP.get_response(DEPLOY_HOOKS_URI)
    end

    redirect_to external_posts_path
  rescue => e
    redirect_back fallback_location: external_posts_path
  end

  def edit; end

  def update
    @external_post.update(allowed_params)
    if Rails.env.production?
      Net::HTTP.get_response(DEPLOY_HOOKS_URI)
    end

    redirect_to external_posts_path
  rescue => e
    redirect_back fallback_location: external_posts_path
  end

  def destroy
    @external_post.destroy!
    if Rails.env.production?
      Net::HTTP.get_response(DEPLOY_HOOKS_URI)
    end

    redirect_to external_posts_path
  rescue => e
    redirect_back fallback_location: external_posts_path
  end


  private

  def find_external_post
    @external_post = ExternalPost.find(params[:id])
  end

  def allowed_params
    params.require(:external_post).permit(
      :title,
      :url,
      :thumbnail_url,
      :category,
      :published_at
    )
  end
end